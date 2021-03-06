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
            <q-select filled v-model="response" :options="responseOptions" label="Please select" @input="filter">
              <template v-if="response" v-slot:append>
                <q-icon name="cancel" @click.stop="resetFilterStatus" class="cursor-pointer" />
              </template>
            </q-select>
          </div>
          <div class="col-md-4">
            <div class="field-name q-mb-xs">Telemarketer</div>
            <q-select
              use-input
              input-debounce="0"
              filled
              v-model="telemarketerName"
              :options="telemarketerNameOptions"
              label="Telemarketer Name"
              @input="filter"
              @filter="filterTelemarketerNameFn"
            >
              <template v-if="telemarketerName" v-slot:append>
                <q-icon name="cancel" @click.stop="resetFilterTelemarketerName" class="cursor-pointer" />
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
        { name: 'userID', label: 'USER ID', align: 'center', field: 'UserID', sortable: true },
        { name: 'name', label: 'NAME', align: 'left', field: 'Name', sortable: true },
        { name: 'phoneNumber', label: 'PHONE NUMBER', align: 'center', field: 'PhoneNumber', sortable: true },
        { name: 'status', label: "STATUS", align: 'center', field: 'Status', sortable: true },
        { name: 'telemarketerName', label: "TELEMARKETER NAME", align: 'center', field: 'TelemarketerName', sortable: true }
      ],
      customerDataFilter: "",
      customerDataVisible: ['userID', 'name', 'phoneNumber', 'status', 'telemarketerName'],
      customerDataPagination: {
        rowsPerPage: 5 // current rows per page being displayed
      },
      telemarketers: [],

      // filter model
      response: null,
      responseOptions: [
        'Tertarik', 'Hubungi Kembali', 'Tidak Tertarik', 'Tidak Aktif', 'Tidak Menjawab', 'Tidak Terdaftar', 'HOT 80%'
      ],
      
      telemarketerName: null,
      telemarketerNameOptionsStr: [],
      telemarketerNameOptions: this.telemarketerNameOptionsStr,
    }
  },

  mounted() {
    this.loadTelemarketer()
  },

  methods: {
    loadTelemarketer() {
      var vm = this;
      var data_submit_telemarketer = {
        Token: vm.$authService.getToken(),
        Limit: 10000,
      }
      this.$axios
      .post("/api/telemarketer/get", data_submit_telemarketer)
      .then(function (response) {
        if (response.data) {
          vm.telemarketers = response.data.Telemarketers
          vm.telemarketerNameOptionsStr = vm.telemarketers.filter(x => x.IsAdmin == false).map(x => x.Name)
          vm.loadCustomer()
        }
      })
    },
    loadCustomer(){
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
            vm.generateTelemarketers()
          }
        })
    },
    generateTelemarketers(){
      var vm = this
      this.customers.forEach(function(data) {
        var index = vm.telemarketers.findIndex(x => x.ID === data.TelemarketerID);
        var telemarketerName = null
        if (index >= 0){
          telemarketerName = vm.telemarketers[index].Name
        }
        vm.$set(data, "TelemarketerName", telemarketerName);
      });
    },
    onRowClick(evt, row) {
      console.log("clicked on", row.PhoneNumber);
      this.$router.push({
        name: "admin-customer-detail",
        params: { phoneNumber: row.PhoneNumber },
      });
    },
    filter() {
      var vm = this;
      var index = vm.telemarketers.findIndex(x => x.Name === vm.telemarketerName);  
      if (vm.telemarketerName == null) {
        var telemarketerID = null
      } else {
        var telemarketerID = "telemarketerID"
      }
      if (index >= 0){
        telemarketerID = vm.telemarketers[index].ID
      }
      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCustomer: {
          Status: vm.response,
          TelemarketerID: telemarketerID,
        },
        Limit: 10000,
      }
      this.$axios
        .post("/api/customer", data_submit)
        .then(function (response) {
          if (response.data) {
            if(response.data.Customers){
              vm.customers = response.data.Customers
              vm.generateTelemarketers()
            }else {
              vm.customers = []
            }
          }
        })
      // console.log(data_submit)
    },
    resetFilterStatus(){
      this.response = null
      this.filter()
    },
    resetFilterTelemarketerName(){
      this.telemarketerName = null
      this.filter()
    },
    filterTelemarketerNameFn(val, update) {
      var vm = this
      if (val === "") {
        update(() => {
          vm.telemarketerNameOptions = vm.telemarketerNameOptionsStr;
        });
        return;
      }
      update(() => {
        const needle = val.toLowerCase();
        vm.telemarketerNameOptions = vm.telemarketerNameOptionsStr.filter(
          (v) => v.toLowerCase().indexOf(needle) > -1
        );
      });
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