<template>
  <q-page class="page-index">
    <div class="row justify-between">
      <div class="title">Customer</div>
      <div>
        <q-btn
          color="light-green-9"
          text-color="white"
          icon="add"
          label="Assign New Customer To Me"
          class="q-pr-sm"
          @click="assignCustomer"
        />
      </div>
    </div>
    <div class="table-filter">
      <q-card class="q-pa-md">
        <q-card-section class="title">
          <div>Filter</div>
        </q-card-section>
        <q-card-section class="row content q-col-gutter-md">
          <div class="col-md-4">
            <div class="field-name q-mb-xs">Status</div>
            <q-select
              filled
              v-model="response"
              :options="options"
              label="Please select"
              @input="filter"
            >
              <template v-if="response" v-slot:append>
                <q-icon name="cancel" @click.stop="resetFilterStatus" class="cursor-pointer" />
              </template>
            </q-select>
          </div>
        </q-card-section>
      </q-card>
    </div>
    <div class="table-container">
      <q-table
        :data="customers"
        :columns="customerDataColumns"
        row-key="phoneNumber"
        :filter="customerDataFilter"
        :visible-columns="customerDataVisible"
        :pagination.sync="customerDataPagination"
        @row-click="onRowClick"
      >
        <template v-slot:top-right>
          <q-input
            color="grey-8"
            dense
            debounce="300"
            v-model="customerDataFilter"
            placeholder="Search"
          >
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>
        <template v-slot:top-left>
          <div>
            <span class="q-pr-sm">Total Customer</span>
            <span class="text-bold" style="font-size: 14px">{{customers.length}}</span>
          </div>
        </template>
        <q-td :style="{width: '170px'}" slot="body-cell-action" slot-scope="props" :props="props">
          <q-btn
            v-if="props.row.Status != ''"
            color="primary"
            icon="book_onlines"
            @click.stop="closing(props.row)"
            label="Closing"
          ></q-btn>
        </q-td>
      </q-table>
    </div>
    <q-dialog v-model="closingPrompt" persistent>
      <q-card style="min-width: 350px">
        <q-form @submit="onClosingSubmit">
          <q-card-section>
            <div class="text-h6">Closing Confirmation | {{closingPhoneNumber}}</div>
          </q-card-section>

          <q-card-section class="content">
            <div class="q-mb-xs">
              <div class="field-name q-mb-xs">
                User ID
                <span style="color: red; font-weight: normal">*</span>
              </div>
              <q-input
                filled
                v-model="closingUserID"
                dense
                :rules="[val => !!val || 'Field is required']"
                @focus="focusInput"
              />
            </div>
            <div class="q-mb-xs">
              <div class="field-name q-mb-xs">
                Name
                <span style="color: red; font-weight: normal">*</span>
              </div>
              <q-input
                filled
                v-model="closingName"
                dense
                :rules="[val => !!val || 'Field is required']"
                @focus="focusInput"
              />
            </div>
            <div class="q-mb-xs">
              <div class="field-name q-mb-xs">
                Deposit Amount
                <span style="color: red; font-weight: normal">*</span>
              </div>
              <q-input
                filled
                v-model="closingBuyAmount"
                dense
                :rules="[val => !!val || 'Field is required']"
                @focus="focusInput"
              >
                <template v-slot:prepend>
                  <div style="font-size: 14px">Rp. </div>
                </template>
              </q-input>
            </div>
            <div v-if="closingFailed" class="text-center" style="font-size: 12px; color: #C10015">
              Closing Failed
            </div>
          </q-card-section>

          <q-card-actions align="right" class="text-primary">
            <q-btn flat label="Cancel" @click="resetClosingPrompt" class="q-px-md"/>
            <q-btn flat label="Ok" type="submit" class="q-px-md"/>
          </q-card-actions>
        </q-form>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script>
export default {
  name: "CustomerIndex",

  data() {
    return {
      customers: [],
      customerDataColumns: [
        {
          name: "name",
          label: "NAME",
          align: "left",
          field: "Name",
          sortable: true,
        },
        {
          name: "phoneNumber",
          label: "PHONE NUMBER",
          align: "center",
          field: "PhoneNumber",
          sortable: true,
        },
        {
          name: "status",
          label: "STATUS",
          align: "center",
          field: "Status",
          sortable: true,
        },
        { name: "action", align: "center", label: "ACTION" },
      ],
      customerDataFilter: "",
      customerDataVisible: ["name", "phoneNumber", "status", "action"],
      customerDataPagination: {
        rowsPerPage: 5, // current rows per page being displayed
      },
      // filter model
      options: [
        "Tertarik",
        "Hubungi Kembali",
        "Tidak Tertarik",
        "Tidak Aktif",
        "Tidak Menjawab",
        "Tidak Terdaftar",
      ],
      response: "",

      closingPrompt: false,
      closingCustomer: {},
      closingPhoneNumber: "",
      closingUserID: "",
      closingName: "",
      closingBuyAmount: "",
      closingFailed: false,
    };
  },

  mounted() {
    this.loadCustomer()
  },

  watch: {
    closingBuyAmount: function(newValue) {
      var vm = this
      const result = newValue.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
      this.$nextTick(() => { vm.closingBuyAmount = result });
    }
  },

  methods: {
    loadCustomer() {
      var vm = this;
      var data_submit = {
        Token: vm.$authService.getToken(),
        Limit: 10000,
      };
      this.$axios.post("/api/customer", data_submit).then(function (response) {
        if (response.data) {
          if (response.data.Customers != null){
            vm.customers = response.data.Customers;
            vm.removeClosedCustomer()
          } else {
            vm.customers = []
          }
        }
      });
    },
    closing(customer) {
      this.closingPrompt = true
      this.closingCustomer = customer
      this.closingPhoneNumber = customer.PhoneNumber
      this.closingName = customer.Name
    },
    onClosingSubmit(){
      var vm = this
      var copyClosingCustomer = this.deepCopyObj(this.closingCustomer)
      copyClosingCustomer.UserID = this.closingUserID
      copyClosingCustomer.Name = this.closingName
      copyClosingCustomer.BuyAmount = parseInt(this.closingBuyAmount.replace(/\./g,''))
      copyClosingCustomer.IsClosing = true
      var data_submit = {
        Token: vm.$authService.getToken(),
        Customer: copyClosingCustomer
      };
      this.$axios
        .post("/api/customer/closing", data_submit)
        .then(function (response) {
          if (response.data) {
            vm.closingPrompt = false
            vm.closingPhoneNumber = ""
            vm.closingUserID = ""
            vm.closingName = ""
            vm.closingBuyAmount = ""
            vm.closingCustomer.IsClosing = true
            vm.removeClosedCustomer()
            vm.closingCustomer = {}
          }
        })
        .catch(function (error) {
          vm.closingFailed = true
        });
    },
    deepCopyObj(src){
      let target = {};
      for (let prop in src) {
          if (src.hasOwnProperty(prop)) {
              target[prop] = src[prop]; //iteratively copies over values, not references
          }
      }
      return target;
    },
    resetClosingPrompt() {
      this.closingPrompt = false
      this.closingCustomer = {}
      this.closingPhoneNumber = ""
      this.closingUserID = ""
      this.closingName = ""
      this.closingBuyAmount = ""
      this.closingFailed = false
    },
    focusInput() {
      this.closingFailed = false  
    },
    assignCustomer() {
      var vm=this;
      var data_submit = {
        Token: vm.$authService.getToken(),
      };
      this.$axios
        .post("/api/customer/assign", data_submit)
        .then(function (response) {
          if (response.data) {
            vm.customers.push(response.data.Customer);
          }
        });
    },
    onRowClick(evt, row) {
      var vm=this;
      this.$router.push({
        name: "customer-detail",
        params: { phoneNumber: row.PhoneNumber },
      });
    },
    filter() {
      var vm = this;
      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCustomer: {
          Status: vm.response,
        },
        Limit: 10000,
      };
      this.$axios.post("/api/customer", data_submit).then(function (response) {
        if (response.data) {
          if (response.data.Customers) {
            vm.customers = response.data.Customers;
            vm.removeClosedCustomer()
          } else {
            vm.customers = [];
          }
        }
      });
      // console.log(data_submit);
    },
    resetFilterStatus(){
      this.response = null
      this.filter()
    },
    removeClosedCustomer() {
      var notClosedCustomer = this.customers.filter(x => x.IsClosing == false);
      this.customers = notClosedCustomer
    },
  },
};
</script>