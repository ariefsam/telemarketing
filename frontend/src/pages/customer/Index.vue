<template>
  <q-page class="page-index">
    <div class="row justify-between">
      <div class="title">Customer</div>
      <!--<div>
        <q-btn
          color="light-green-9"
          text-color="white"
          icon="add"
          label="Request"
          class="q-pr-sm"
        />
      </div>-->
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
            <div class="col-md-4">
              <br />
              <q-btn @click="assignCustomer" label="Assign New Customer To Me"></q-btn>
            </div>
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
            color="primary"
            icon="book_onlines"
            @click.stop="closing(props.row)"
            label="Closing"
          ></q-btn>
        </q-td>
      </q-table>
    </div>
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
    };
  },

  mounted() {
    this.loadCustomer()
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
      var vm = this;
      vm.$q
        .dialog({
          title: "Closing Confirmation | " + customer.PhoneNumber,
          message: "Customer Name",
          prompt: {
            model: customer.Name,
            isValid: val => val.length > 3,
            type: 'text' // optional
          },
          cancel: true,
          persistent: true,
        })
        .onOk(data => {
          customer.Name = data
          customer.IsClosing = true
          var data_submit = {
            Token: vm.$authService.getToken(),
            Customer: customer
          };
          console.log(customer.IsClosing)
          this.$axios
            .post("/api/customer/save", data_submit)
            .then(function (response) {
              if (response.data) {
                vm.removeClosedCustomer()
              }
            })
            .catch(function (error) {
              alert("Closing Failed")
            });
        })
        .onOk(() => {
          // console.log('>>>> second OK catcher')
        })
        .onCancel(() => {
          // console.log('>>>> Cancel')
        })
        .onDismiss(() => {
          // console.log('I am triggered on both OK and Cancel')
        });
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