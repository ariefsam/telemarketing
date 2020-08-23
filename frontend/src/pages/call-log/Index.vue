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
          <div class="col-md-12">
            <div class="row">
              <q-input
                filled
                color="grey-8"
                v-model="start_date"
                label="Start Date *"
                hint="Start date"
                class="col-sm-6 col-xs-12 q-pr-sm"
              >
                <template v-slot:append>
                  <q-icon name="event" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-date v-model="start_date" mask="DD-MM-YYYY" />
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
              <q-input
                filled
                color="grey-8"
                v-model="start_time"
                label="Start Time *"
                hint="Start time"
                class="col-sm-6 col-xs-12 q-pl-sm"
              >
                <template v-slot:append>
                  <q-icon name="access_time" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-time v-model="start_time" mask="HH:mm" format24h />
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
            </div>
            <div class="row">
              <q-input
                filled
                color="grey-8"
                v-model="end_date"
                label="End Date *"
                hint="End date"
                class="col-sm-6 col-xs-12 q-pr-sm"
              >
                <template v-slot:append>
                  <q-icon name="event" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-date v-model="end_date" mask="DD-MM-YYYY" />
                    </q-popup-proxy>
                  </q-icon>
                </template>
              </q-input>
              <q-input
                filled
                color="grey-8"
                v-model="end_time"
                label="End Time *"
                hint="End time"
                class="col-sm-6 col-xs-12 q-pl-sm"
              >
                <template v-slot:append>
                  <q-icon name="access_time" class="cursor-pointer">
                    <q-popup-proxy transition-show="scale" transition-hide="scale">
                      <q-time v-model="end_time" mask="HH:mm" format24h />
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
          <div>
            <span class="q-pr-sm">Total Call Logs</span>
            <span class="text-bold" style="font-size: 14px">{{data.length}}</span>
          </div>
        </template>
        <q-td :style="{width: '170px'}" slot="body-cell-action" slot-scope="props" :props="props">
          <q-btn
            color="primary"
            icon="book_onlines"
            class="action-btn q-mx-xs"
            @click.stop="closing(props.row.PhoneNumber)"
            label="Closing"
          ></q-btn>
        </q-td>
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
        { name: "action", align: "center", label: "Action" },
      ],
      dataFilter: "",
      dataVisible: ["date", "name", "phoneNumber", "status", "action"],
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
    };
  },

  mounted() {
    this.filter();
    this.start_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
    this.end_date = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
  },

  methods: {
    closing(phoneNumber) {
      var vm = this;
      vm.$q
        .dialog({
          title: "Closing Confirmation",
          message: "Yakin untuk closing?",
          cancel: true,
          persistent: true,
        })
        .onOk(() => {
          var data_submit = {
            Token: vm.$authService.getToken(),
            PhoneNumber: phoneNumber,
            Status: "Closing",
          };
          vm.$restapi.callCustomer(data_submit,()=>{

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