<template>
  <q-page class="page-index">
    <div class="row justify-between">
      <div class="title">Call Log</div>
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
          <div class="col-md-12">
            <div class="field-name q-mb-xs">Status</div>
            <q-select
              filled
              v-model="response"
              :options="options"
              label="Please select"
              @input="filter"
            />
          </div>
          <div class="col-md-6">
            <div class="field-name q-mb-xs">Timestamp</div>
            <q-select
              filled
              v-model="filterTimestamp"
              :options="filterTimestampOptions"
              label="Please select"
              @input="filter"
            >
              <template v-if="filterTimestamp" v-slot:append>
                <q-icon name="cancel" @click.stop="clearFilterTimestamp" class="cursor-pointer" />
              </template>
            </q-select>
          </div>
          <div class="col-md-12">
            <div class="field-name q-mb-xs">Custom Timestamp</div>
            <div class="row">
              <q-input
                filled
                color="grey-8"
                v-model="start_date"
                label="Start Date"
                class="col-sm-6 col-xs-12 q-pr-xs q-mb-sm"
                @keydown.enter.prevent="inputStartDate"
              >
                <template v-slot:append>
                  <q-icon name="event" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-date v-model="start_date" mask="DD-MM-YYYY" @input="inputStartDate"/>
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
              <q-input
                filled
                color="grey-8"
                v-model="end_date"
                label="End Date"
                class="col-sm-6 col-xs-12 q-pl-xs q-mb-sm"
                @keydown.enter.prevent="inputEndDate"
              >
                <template v-slot:append>
                  <q-icon name="event" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-date v-model="end_date" mask="DD-MM-YYYY" @input="inputEndDate"/>
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
            </div>
            <div class="row">
              <q-input
                filled
                color="grey-8"
                v-model="start_time"
                label="Start Time"
                class="col-sm-6 col-xs-12 q-pr-xs"
                @keydown.enter.prevent="inputStartTime"
              >
                <template v-slot:append>
                  <q-icon name="access_time" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-time v-model="start_time" mask="HH:mm" format24h @input="inputStartTime"/>
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
              <q-input
                filled
                color="grey-8"
                v-model="end_time"
                label="End Time"
                class="col-sm-6 col-xs-12 q-pl-xs"
                @keydown.enter.prevent="inputEndTime"
              >
                <template v-slot:append>
                  <q-icon name="access_time" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-time v-model="end_time" mask="HH:mm" format24h @input="inputEndTime"/>
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
            </div>
          </div>
        </q-card-section>
      </q-card>
    </div>
    <div class="table-container">
      <q-table
        :data="data"
        :columns="dataColumns"
        row-key="Timestamp"
        :filter="dataFilter"
        :visible-columns="dataVisible"
        :pagination.sync="dataPagination"
        @row-click="onRowClick"
      >
        <template v-slot:top-right>
          <q-input color="grey-8" dense debounce="300" v-model="dataFilter" placeholder="Search">
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>
        <template v-slot:top-left>
          <div class="q-mb-sm">
            <span class="q-pr-sm">Total Call Logs</span>
            <span class="text-bold" style="font-size: 14px">{{data.length}}</span>
          </div>
          <div class="row justify-between">
            <div class="column q-mr-md">
              <div>Total Tertarik</div>
              <div>Total Hubungi Kembali</div>
              <div>Total Tidak Tertarik</div>
              <div>Total Tidak Aktif</div>
              <div>Total Tidak Menjawab</div>
              <div>Total Tidak Terdaftar</div>
            </div>
            <div class="column">
              <div class="text-bold">{{totalTertarik}}</div>
              <div class="text-bold">{{totalHubungiKembali}}</div>
              <div class="text-bold">{{totalTidakTertarik}}</div>
              <div class="text-bold">{{totalTidakAktif}}</div>
              <div class="text-bold">{{totalTidakMenjawab}}</div>
              <div class="text-bold">{{totalTidakTerdaftar}}</div>
            </div>
          </div>
        </template>
      </q-table>
    </div>
  </q-page>
</template>

<script>
import { date } from "quasar";
import moment from "moment-timezone";

export default {
  name: "CustomerIndex",

  data() {
    return {
      start_date: "",
      start_time: "00:00",
      end_date: "",
      end_time: "23:59",
      data: [],
      dataColumns: [
        {
          name: "date",
          label: "DATE",
          align: "left",
          field: "Date",
          sortable: true,
        },
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
      ],
      dataFilter: "",
      dataVisible: ["date", "name", "phoneNumber", "status"],
      dataPagination: {
        rowsPerPage: 5, // current rows per page being displayed
      },
      // filter model
      options: [
        "All",
        "Tertarik",
        "Hubungi Kembali",
        "Tidak Tertarik",
        "Tidak Aktif",
        "Tidak Menjawab",
        "Tidak Terdaftar",
      ],
      response: "All",
      totalTertarik: 0,
      totalHubungiKembali: 0,
      totalTidakTertarik: 0,
      totalTidakAktif: 0,
      totalTidakMenjawab: 0,
      totalTidakTerdaftar: 0,
      filterTimestamp: "",
      filterTimestampOptions: [
        "Today",
        "Yesterday",
        "Last 7 days",
        "Last 15 days",
        "Last 30 days",
        "Last 60 days",
        "Last 90 days",
      ],
    };
  },

  mounted() {
    this.start_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
    this.end_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
    this.filter();
  },

  methods: {
    assignDataFromAPI(dataResponse) {
      var vm = this;
      var data = dataResponse;
      Object.keys(data).forEach((key) => {
        var Timestamp = new Date(data[key].Timestamp / 1000000);
        Timestamp = date.formatDate(Timestamp, "DD-MM-YYYY HH:mm:ss Z");

        var dataModel = data[key];
        dataModel.Date = Timestamp;

        vm.$set(vm.data, key, dataModel);
      });
      vm.totalTertarik = vm.data.filter(x => x.Status == "Tertarik").length
      vm.totalHubungiKembali = vm.data.filter(x => x.Status == "Hubungi Kembali").length
      vm.totalTidakTertarik = vm.data.filter(x => x.Status == "Tidak Tertarik").length
      vm.totalTidakAktif = vm.data.filter(x => x.Status == "Tidak Aktif").length
      vm.totalTidakMenjawab = vm.data.filter(x => x.Status == "Tidak Menjawab").length
      vm.totalTidakTerdaftar = vm.data.filter(x => x.Status == "Tidak Terdaftar").length
    },
    onRowClick(evt, row) {
      var vm = this;
      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCustomer: {
          PhoneNumber: row.PhoneNumber,
        },
        Limit: 10000,
      };
      var isThisCustomerClosed = false
      this.$axios.post("/api/customer", data_submit).then(function (response) {
        if (response.data) {
          isThisCustomerClosed = response.data.Customers[0].IsClosing;
          if (isThisCustomerClosed){
            vm.$q
              .dialog({
                title: "Info",
                message: "Customer : "+row.Name+" ("+row.PhoneNumber+") is already Closed",
                persistent: false,
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
          }else {
            vm.$router.push({
              name: "customer-detail",
              params: { phoneNumber: row.PhoneNumber },
            });
          }
        }
      });
    },
    optionsStartdate (date) {
      // return moment(String(date)).format('DD-MM-YYYY') <= this.end_date
    },
    optionsEnddate (date) {
      // return moment(String(date)).format('DD-MM-YYYY') >= this.start_date
    },
    getFromToTimestamp(filterTimestamp) {
      console.log(filterTimestamp)
      var StartEndDate = {
        from: "",
        to: ""
      }
      if (this.filterTimestamp == "Today") {
        StartEndDate.from = moment().startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (this.filterTimestamp == "Yesterday") {
        StartEndDate.from = moment().subtract(1, "days").startOf("day");
        StartEndDate.to = moment().subtract(1, "days").endOf("day");
      } else if (this.filterTimestamp == "Last 7 days") {
        StartEndDate.from = moment().subtract(7, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (this.filterTimestamp == "Last 15 days") {
        StartEndDate.from = moment().subtract(15, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (this.filterTimestamp == "Last 30 days") {
        StartEndDate.from = moment().subtract(30, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (this.filterTimestamp == "Last 60 days") {
        StartEndDate.from = moment().subtract(60, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (this.filterTimestamp == "Last 90 days") {
        StartEndDate.from = moment().subtract(90, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      }
      // console.log(StartEndDate.from.format("DD-MM-YYYY HH:mm"))
      // console.log(StartEndDate.to.format("DD-MM-YYYY HH:mm"))
      return StartEndDate
    },
    filter() {
      var vm = this;
      vm.data = [];
      vm.totalTertarik = 0
      vm.totalHubungiKembali = 0
      vm.totalTidakTertarik = 0
      vm.totalTidakAktif = 0
      vm.totalTidakMenjawab = 0
      vm.totalTidakTerdaftar = 0

      var fromDateTime = ""
      var toDateTime = ""
      if (this.filterTimestamp) {
        var StartEndDate = this.getFromToTimestamp(this.filterTimestamp)
        fromDateTime = StartEndDate.from
        toDateTime = StartEndDate.to
        this.start_date = fromDateTime.format("DD-MM-YYYY")
        this.end_date = toDateTime.format("DD-MM-YYYY")
        this.start_time = "00:00"
        this.end_time = "23:59"
      } else {
        var start_datetime = this.start_date + " " + this.start_time;
        var end_datetime = this.end_date + " " + this.end_time;
        fromDateTime = moment.tz(start_datetime, "DD-MM-YYYY HH:mm", "Asia/Jakarta");
        toDateTime = moment.tz(end_datetime, "DD-MM-YYYY HH:mm", "Asia/Jakarta");
      }
      
      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCallLog: {
          Status: vm.response,
          TimestampStart: fromDateTime.unix() * 1000000000,
          TimestampEnd: toDateTime.unix() * 1000000000,
        },
        Limit: 10000,
      };

      if (this.response == "All") {
        delete data_submit.FilterCallLog.Status;
      }
      this.$axios
        .post("/api/call-log/get", data_submit)
        .then(function (response) {
          if (response.data) {
            if (response.data.CallLogs) {
              vm.assignDataFromAPI(response.data.CallLogs);
            } else {
              vm.data = [];
            }
          }
        });
      // console.log(data_submit);
    },
    clearFilterTimestamp() {
      this.filterTimestamp = "";
      this.start_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
      this.end_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
      this.start_time = "00:00"
      this.end_time = "23:59"
      this.filter();
    },
    inputStartDate() {
      this.filterTimestamp = "";
      if(this.start_date == ""){
        this.start_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
      }
      this.filter();
    },
    inputEndDate() {
      this.filterTimestamp = "";
      if(this.end_date == ""){
        this.end_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
      }
      this.filter();
    },
    inputStartTime() {
      this.filterTimestamp = "";
      if(this.start_time == ""){
        this.start_time = "00:00"
      }
      this.filter();
    },
    inputEndTime() {
      this.filterTimestamp = "";
      if(this.end_time == ""){
        this.end_time = "23:59"
      }
      this.filter();
    },
  },
};
</script>