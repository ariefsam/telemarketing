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
          <div class="q-gutter-sm">
            <q-radio dense v-model="period" val="today" label="Today" />
            <q-radio dense v-model="period" val="yesterday" label="Yesterday" />
            <q-radio dense v-model="period" val="thisWeek" label="This Week" />
            <q-radio dense v-model="period" val="last Week" label="Last Week" />
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
          <div>
            <span class="q-pr-sm">Total Call Logs</span>
            <span class="text-bold" style="font-size: 14px">{{data.length}}</span>
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
      period: "",
      start_date: "",
      start_time: "00:00",
      end_date: "",
      end_time: "23:59",
      data: [],
      dataColumns: [
        {
          name: "date",
          label: "Date",
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
        "HOT 80%",
      ],
      response: "",
    };
  },

  mounted() {
    this.filter();
    this.start_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
    this.end_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
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
        console.log(dataModel);
      });
    },
    onRowClick(evt, row) {
      var vm = this;
      // console.log("clicked on", row.PhoneNumber);
      this.$router.push({
        name: "customer-detail",
        params: { phoneNumber: row.PhoneNumber },
      });
    },
    filter() {
      var vm = this;
      vm.data = [];

      var start_datetime = this.start_date + " " + this.start_time;
      var end_datetime = this.end_date + " " + this.end_time;
      if (this.is_private) {
        var list_of_audience = this.invited_audiences
          .split("\n")
          .filter((v) => v != "");
      }

      var startMoment = moment.tz(
        start_datetime,
        "DD-MM-YYYY HH:mm",
        "Asia/Jakarta"
      );
      var endMoment = moment.tz(
        end_datetime,
        "DD-MM-YYYY HH:mm",
        "Asia/Jakarta"
      );

      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCallLog: {
          Status: vm.response,
          TimestampStart: startMoment.unix() * 1000000000,
          TimestampEnd: endMoment.unix() * 1000000000,
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
  },
};
</script>