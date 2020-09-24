<template>
  <q-page class="page-index">
    <div class="row justify-between">
      <div class="title">Closed Customer</div>
    </div>
    <!--<div class="table-filter">
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
    </div>-->
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
      </q-table>
    </div>
  </q-page>
</template>

<script>
import moment from "moment-timezone";
import { date } from 'quasar'

export default {
  name: "CustomerClosed",

  data() {
    return {
      customers: [],
      customerDataColumns: [
        { name: 'userID', label: 'USER ID', align: 'center', field: 'UserID', sortable: true },
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
          name: "buyAmount",
          label: "DEPOSIT AMOUNT",
          align: "center",
          field: "BuyAmount",
          sortable: true,
          format: (val, row) => `Rp. ${this.moneyFormat(val)}`,
        },
        {
          name: "closingDate",
          label: "CLOSING DATE",
          align: "center",
          field: "ClosingTimestamp",
          sortable: true,
          format: (val, row) => `${this.timestampFormat(val)}`,
        },
      ],
      customerDataFilter: "",
      customerDataVisible: ["userID", "name", "phoneNumber", "buyAmount", "closingDate"],
      customerDataPagination: {
        rowsPerPage: 5, // current rows per page being displayed
      },
      // // filter model
      // options: [
      //   "Tertarik",
      //   "Hubungi Kembali",
      //   "Tidak Tertarik",
      //   "Tidak Aktif",
      //   "Tidak Menjawab",
      //   "Tidak Terdaftar",
      // ],
      // response: "",
    };
  },

  mounted() {
    this.loadCustomer()
  },

  methods: {
    moneyFormat(value){
      const result = value.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".")
      return result
    },
    timestampFormat(value){
      var result = new Date(value / 1000000)
      result = date.formatDate(result, 'DD-MM-YYYY')
      return result
    },
    loadCustomer() {
      var vm = this;
      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCustomer: {
          IsClosing: true
        },
        Limit: 10000,
      };
      this.$axios.post("/api/customer", data_submit).then(function (response) {
        if (response.data) {
          if (response.data.Customers != null){
            vm.customers = response.data.Customers;
          } else {
            vm.customers = []
          }
        }
      });
    },
    onRowClick(evt, row) {
      // this.$router.push({
      //   name: "customer-detail",
      //   params: { phoneNumber: row.PhoneNumber },
      // });
    },
    // filter() {
    //   var vm = this;
    //   var data_submit = {
    //     Token: vm.$authService.getToken(),
    //     FilterCustomer: {
    //       Status: vm.response,
    //     },
    //     Limit: 10000,
    //   };
    //   this.$axios.post("/api/customer", data_submit).then(function (response) {
    //     if (response.data) {
    //       if (response.data.Customers) {
    //         vm.customers = response.data.Customers;
    //       } else {
    //         vm.customers = [];
    //       }
    //     }
    //   });
    //   // console.log(data_submit);
    // },
    // resetFilterStatus(){
    //   this.response = null
    //   this.filter()
    // },
  },
};
</script>