<template>
  <q-page class="page-index">
    <div class="row justify-between">
      <div class="title">Customer List</div>
      <div>
        <q-btn
          color="light-green-9"
          text-color="white"
          icon="fas fa-file-import"
          label="Import"
          class="q-pr-sm"
          @click="importCustomer"
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
            <q-select filled v-model="response" :options="responseOptions" label="Please select" @input="filter"/>
          </div>
          <div class="col-md-4">
            <div class="field-name q-mb-xs">Telemarketer</div>
            <q-input filled v-model="telemarketer" label="Telemarketer email" @keydown.enter.prevent="filter"/>
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
          <q-input color="grey-8" dense debounce="300" v-model="customerDataFilter" placeholder="Search">
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
      </q-table>
    </div>
  </q-page>
</template>

<script>
export default {
  name: 'AdminCustomerIndex',

  data() {
    return {
      customers: [],
      customerDataColumns: [
        { name: 'name', label: 'NAME', align: 'left', field: 'Name', sortable: true },
        { name: 'phoneNumber', label: 'PHONE NUMBER', align: 'center', field: 'PhoneNumber', sortable: true },
        { name: 'status', label: "STATUS", align: 'center', field: 'Status', sortable: true },
        { name: 'telemarketer', label: "TELEMARKETER", align: 'left', field: 'TelemarketerEmail', sortable: true }
      ],
      customerDataFilter: "",
      customerDataVisible: ['name', 'phoneNumber', 'status', 'telemarketer'],
      customerDataPagination: {
        rowsPerPage: 5 // current rows per page being displayed
      },
      // filter model
      responseOptions: [
        'No Status', 'Tertarik', 'Hubungi Kembali', 'Tidak Tertarik', 'Tidak Aktif', 'Tidak Menjawab', 'Tidak Terdaftar'
      ],
      response: "",
      telemarketer: "",
    }
  },

  mounted() {
    var vm=this;
    var data_submit = {
      Token: vm.$authService.getToken(),
      Limit: 10000,
    }
    this.$axios
      .post("/api/customer", data_submit)
      .then(function (response) {
        if (response.data) {
          vm.customers = response.data.Customers
        }
      })
  },

  methods: {
    onRowClick(evt, row) {
      console.log("clicked on", row.PhoneNumber);
      this.$router.push({
        name: "customer-detail",
        params: { phoneNumber: row.PhoneNumber },
      });
    },
    filter() {
      if(this.response == 'No Status'){
        this.response = ''
      }
      var vm=this;
      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCustomer: {
          Status: vm.response,
          TelemarketerEmail: vm.telemarketer,
        },
        Limit: 10000,
      }
      this.$axios
        .post("/api/customer", data_submit)
        .then(function (response) {
          if (response.data) {
            if(response.data.Customers){
              vm.customers = response.data.Customers
            }else {
              vm.customers = []
            }
          }
        })
      console.log(data_submit)
    },
    importCustomer() {
      var vm=this;
      this.$q.dialog({
        title: 'Data Source',
        message: 'What is the data source?',
        prompt: {
          model: '',
          isValid: val => val.length > 2,
          type: 'text' // optional

        },
        cancel: true,
        persistent: true
      }).onOk(data => {
        vm.$router.push({
          name: 'import-customer', params: { source: data },
        });
      }).onCancel(() => {
        // console.log('>>>> Cancel')
      }).onDismiss(() => {
        // console.log('I am triggered on both OK and Cancel')
      })
    }
  }
}
</script>