<template>
  <q-page class="page-index">
    <div class="title">Dashboard</div>
    <div class="row q-col-gutter-md telemarketer-dashboard">
      <div class="col-md-4">
        <div class="section">
          <div class="title-section text-center">Revenue</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary">
                  My {{selectedRevenue}} Target Revenue
                </div>
                <div>
                  <q-select outlined v-model="selectedRevenue" :options="revenueOptions" dense style="min-width: 100px" @input="processingRevenueData"/>
                </div>
              </div>
              <div class="q-mb-sm">
                <div class="rev-performance">
                  <span style="font-size: 24px;">Rp.</span> {{revenue.performanceStr}}
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="revenue.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between q-mb-md rev-target">
                <div>
                  from target <span style="font-size:15px; font-weight: bold">Rp. {{revenue.targetStr}}</span>
                </div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(revenue.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div class="col-md-4">
        <div class="section">
          <div class="title-section text-center">Closing</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary">
                  My {{selectedClosing}} Target Closing
                </div>
                <div>
                  <q-select outlined v-model="selectedClosing" :options="closingOptions" dense style="min-width: 100px" @input="processingClosingData"/>
                </div>
              </div>
              <div class="q-mb-sm">
                <div class="rev-performance row justify-end">
                  <div>{{closing.performance}} <span style="font-size: 16px;color: grey">/ {{closing.target}} Closings</span></div>
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="closing.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between q-mb-md rev-target">
                <div></div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(closing.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div class="col-md-4">
        <div class="section">
          <div class="title-section text-center">Call</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary">
                  My {{selectedCall}} Target Call
                </div>
                <div>
                  <q-select outlined v-model="selectedCall" :options="callOptions" dense style="min-width: 100px" @input="processingCallData"/>
                </div>
              </div>
              <div class="q-mb-sm">
                <div class="rev-performance row justify-end">
                  <div>{{call.performance}} <span style="font-size: 16px;color: grey">/ {{call.target}} Calls</span></div>
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="call.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between q-mb-md rev-target">
                <div></div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(call.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div class="col-md-6">
        <div class="section">
          <div class="title-section text-center">Call Log</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary">
                  My Call Log
                </div>
                <div>
                  <q-select outlined v-model="selectedCallLog" :options="callLogOptions" dense style="min-width: 100px" @input="processingCallLog"/>
                </div>
              </div>
              <div>
                <CallLogChart />
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script>
import moment from "moment-timezone";
import CallLogChart from "components/CallLogChart.vue";

export default {
  name: 'DashboardTelemarketer',

  components: {
    CallLogChart,
  },

  data() {
    return {
      telemarketer: {
        Performance: {
          Daily: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
            BuyAmountStr: "0",
          },
          Weekly: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
            BuyAmountStr: "0",
          },
          Monthly: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
            BuyAmountStr: "0",
          },
        },
        Target: {
          Daily: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
            BuyAmountStr: "0",
          },
          Weekly: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
            BuyAmountStr: "0",
          },
          Monthly: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
            BuyAmountStr: "0",
          },
        }
      },
      
      selectedRevenue: "Daily",
      revenueOptions: ["Daily", "Weekly", "Monthly"],
      revenue: {
        performanceStr: "0",
        targetStr: "0",
        progress: 0,
        progressPercentage: 0,
      },

      selectedClosing: "Daily",
      closingOptions: ["Daily", "Weekly", "Monthly"],
      closing: {
        performance: 0,
        target: 0,
        progress: 0,
        progressPercentage: 0,
      },

      selectedCall: "Daily",
      callOptions: ["Daily", "Weekly", "Monthly"],
      call: {
        performance: 0,
        target: 0,
        progress: 0,
        progressPercentage: 0,
      },

      callLog: [],
      totalTertarik: 0,
      totalHubungiKembali: 0,
      totalTidakTertarik: 0,
      totalTidakAktif: 0,
      totalTidakMenjawab: 0,
      totalTidakTerdaftar: 0,
      selectedCallLog: "Today",
      callLogOptions: [
        "Today",
        "Yesterday",
        "Last 7 days",
        "Last 15 days",
        "Last 30 days",
        "Last 60 days",
        "Last 90 days",
      ],

    }
  },

  mounted() {
    var vm = this
    var user = this.$authService.getUser();
    var data_submit = {
      Token: vm.$authService.getToken(),
      FilterTelemarketer: {
        ID: user.ID,
      },
      Limit: 10000,
    }
    this.$axios
      .post("/api/me", data_submit)
      .then(function (response) {
        if (response.data) {
          vm.telemarketer = response.data.Telemarketers[0]
          vm.$nextTick(() => { 
            vm.processingRevenueData() 
            vm.processingClosingData()
            vm.processingCallData()
          })
        }
      })
    this.processingCallLog();
  },

  watch: {
    telemarketer: {
      handler(val){
        var vm = this
        const resultDailyPerformance = val.Performance.Daily.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultWeeklyPerformance = val.Performance.Weekly.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultMonthlyPerformance = val.Performance.Monthly.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultDailyTarget = val.Target.Daily.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultWeeklyTarget = val.Target.Weekly.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultMonthlyTarget = val.Target.Monthly.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        vm.telemarketer.Performance.Daily.BuyAmountStr = resultDailyPerformance
        vm.telemarketer.Performance.Weekly.BuyAmountStr = resultWeeklyPerformance 
        vm.telemarketer.Performance.Monthly.BuyAmountStr = resultMonthlyPerformance 
        vm.telemarketer.Target.Daily.BuyAmountStr = resultDailyTarget
        vm.telemarketer.Target.Weekly.BuyAmountStr = resultWeeklyTarget 
        vm.telemarketer.Target.Monthly.BuyAmountStr = resultMonthlyTarget 
      },
      deep: true
    }
  },

  methods: {
    processingRevenueData() {
      if (this.selectedRevenue == "Daily") {
        this.revenue.performanceStr = this.telemarketer.Performance.Daily.BuyAmountStr
        this.revenue.targetStr = this.telemarketer.Target.Daily.BuyAmountStr
        this.revenue.progress =  this.telemarketer.Performance.Daily.BuyAmount/this.telemarketer.Target.Daily.BuyAmount
      } else if (this.selectedRevenue == "Weekly"){
        this.revenue.performanceStr = this.telemarketer.Performance.Weekly.BuyAmountStr
        this.revenue.targetStr = this.telemarketer.Target.Weekly.BuyAmountStr
        this.revenue.progress =  this.telemarketer.Performance.Weekly.BuyAmount/this.telemarketer.Target.Weekly.BuyAmount
      } else if (this.selectedRevenue == "Monthly"){
        this.revenue.performanceStr = this.telemarketer.Performance.Monthly.BuyAmountStr
        this.revenue.targetStr = this.telemarketer.Target.Monthly.BuyAmountStr
        this.revenue.progress =  this.telemarketer.Performance.Monthly.BuyAmount/this.telemarketer.Target.Monthly.BuyAmount
      }
      this.$nextTick(() => { 
      Number.isNan(this.revenue.progress) ? 0 : this.revenue.progress
      })
    },
    processingClosingData() {
      if (this.selectedClosing == "Daily") {
        this.closing.performance = this.telemarketer.Performance.Daily.Closing
        this.closing.target = this.telemarketer.Target.Daily.Closing
        this.closing.progress =  this.telemarketer.Performance.Daily.Closing/this.telemarketer.Target.Daily.Closing
      } else if (this.selectedClosing == "Weekly"){
        this.closing.performance = this.telemarketer.Performance.Weekly.Closing
        this.closing.target = this.telemarketer.Target.Weekly.Closing
        this.closing.progress =  this.telemarketer.Performance.Weekly.Closing/this.telemarketer.Target.Weekly.Closing
      } else if (this.selectedClosing == "Monthly"){
        this.closing.performance = this.telemarketer.Performance.Monthly.Closing
        this.closing.target = this.telemarketer.Target.Monthly.Closing
        this.closing.progress =  this.telemarketer.Performance.Monthly.Closing/this.telemarketer.Target.Monthly.Closing
      }
    },
    processingCallData() {
      if (this.selectedCall == "Daily") {
        this.call.performance = this.telemarketer.Performance.Daily.Call
        this.call.target = this.telemarketer.Target.Daily.Call
        this.call.progress =  this.telemarketer.Performance.Daily.Call/this.telemarketer.Target.Daily.Call
      } else if (this.selectedCall == "Weekly"){
        this.call.performance = this.telemarketer.Performance.Weekly.Call
        this.call.target = this.telemarketer.Target.Weekly.Call
        this.call.progress =  this.telemarketer.Performance.Weekly.Call/this.telemarketer.Target.Weekly.Call
      } else if (this.selectedCall == "Monthly"){
        this.call.performance = this.telemarketer.Performance.Monthly.Call
        this.call.target = this.telemarketer.Target.Monthly.Call
        this.call.progress =  this.telemarketer.Performance.Monthly.Call/this.telemarketer.Target.Monthly.Call
      }
    },
    getFromToTimestamp(selectedCallLog) {
      var StartEndDate = {
        from: "",
        to: ""
      }
      if (selectedCallLog == "Today") {
        StartEndDate.from = moment().startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (selectedCallLog == "Yesterday") {
        StartEndDate.from = moment().subtract(1, "days").startOf("day");
        StartEndDate.to = moment().subtract(1, "days").endOf("day");
      } else if (selectedCallLog == "Last 7 days") {
        StartEndDate.from = moment().subtract(7, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (selectedCallLog == "Last 15 days") {
        StartEndDate.from = moment().subtract(15, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (selectedCallLog == "Last 30 days") {
        StartEndDate.from = moment().subtract(30, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (selectedCallLog == "Last 60 days") {
        StartEndDate.from = moment().subtract(60, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (selectedCallLog == "Last 90 days") {
        StartEndDate.from = moment().subtract(90, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      }
      return StartEndDate
    },
    processingCallLog() {
      var vm = this
      vm.callLog = [];
      vm.totalTertarik = 0
      vm.totalHubungiKembali = 0
      vm.totalTidakTertarik = 0
      vm.totalTidakAktif = 0
      vm.totalTidakMenjawab = 0
      vm.totalTidakTerdaftar = 0

      var fromDateTime = ""
      var toDateTime = ""
      var StartEndDate = this.getFromToTimestamp(this.selectedCallLog)
      fromDateTime = StartEndDate.from
      toDateTime = StartEndDate.to

      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterCallLog: {
          TimestampStart: fromDateTime.unix() * 1000000000,
          TimestampEnd: toDateTime.unix() * 1000000000,
        },
        Limit: 10000,
      };

      this.$axios
        .post("/api/call-log/get", data_submit)
        .then(function (response) {
          if (response.data) {
            if (response.data.CallLogs) {
              vm.assignDataCallLog(response.data.CallLogs);
            } else {
              console.log("tidak ada")
            }
          }
        });
    },
    assignDataCallLog(dataResponse) {
      this.totalTertarik = dataResponse.filter(x => x.Status == "Tertarik").length
      this.totalHubungiKembali = dataResponse.filter(x => x.Status == "Hubungi Kembali").length
      this.totalTidakTertarik = dataResponse.filter(x => x.Status == "Tidak Tertarik").length
      this.totalTidakAktif = dataResponse.filter(x => x.Status == "Tidak Aktif").length
      this.totalTidakMenjawab = dataResponse.filter(x => x.Status == "Tidak Menjawab").length
      this.totalTidakTerdaftar = dataResponse.filter(x => x.Status == "Tidak Terdaftar").length
    },
  }
}
</script>
