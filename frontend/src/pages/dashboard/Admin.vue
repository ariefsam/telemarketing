<template>
  <q-page class="page-index">
    <div class="title">Dashboard</div>
    <div class="row q-col-gutter-md admin-dashboard">
      <div class="col-md-12">
        <div class="section">
          <div class="title-section text-center">Telemarketer Performance Filter</div>
          <q-card>
            <q-card-section class="row">
              <div class="col-md-4">
                <div class="field-name q-mb-xs">Custom Timestamp</div>
                <q-input
                  outlined
                  v-model="filterCustomTimestamp"
                  dense
                  @keydown.enter.prevent="inputFilterCustomTimestamp"
                >
                  <template v-slot:append>
                    <q-icon name="event" class="cursor-pointer">
                      <q-popup-proxy transition-show="scale" transition-hide="scale">
                        <q-date v-model="filterCustomTimestamp" mask="DD-MMM-YYYY" @input="inputFilterCustomTimestamp"/>
                      </q-popup-proxy>
                    </q-icon>
                  </template>
                </q-input>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </div>
    <div class="row q-col-gutter-md admin-dashboard">
      <div class="col-md-8">
        <div class="section">
          <div class="title-section text-center">Telemarketer Performance</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary q-mb-sm">
                  Telemarketer {{selectedTeleTimestamp}} {{selectedTeleType}}
                </div>
                <div class="row no-wrap">
                  <div class="q-mr-sm">
                    <q-select outlined v-model="selectedTeleTimestamp" :options="teleTimestampOptions" dense style="min-width: 120px" @input="processingTeleData"/>
                  </div>
                  <div>
                    <q-select outlined v-model="selectedTeleType" :options="teleTypeOptions" dense style="min-width: 120px" @input="processingTeleData"/>
                  </div>
                </div>
              </div>
              <div class="telemarketer-chart">
                <TelemarketerChart v-bind:teleChartData="teleChartData" v-if="teleChartData.label.length > 0"/>
                <q-spinner-pie
                  class="loading"
                  color="primary"
                  size="50px"
                  v-if="telemarketerChartLoading"
                />
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div class="col-md-4">
        <div class="section">
          <div class="title-section text-center">Summary</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary">
                  {{selectedTotalTeleTimestamp}} Performance
                </div>
                <div class="row no-wrap">
                  <div class="q-mr-sm">
                    <q-select outlined v-model="selectedTotalTeleTimestamp" :options="totalTeleTimestampOptions" dense style="min-width: 100px" @input="processingTotalTeleData"/>
                  </div>
                </div>
              </div>
              <div class="title-subsummary q-mb-sm text-center">
                Total {{selectedTotalTeleTimestamp}} Target Revenue
              </div>
              <div class="q-mb-md">
                <div class="rev-performance">
                  <span style="font-size: 24px;">Rp.</span> {{revenue.performanceStr}}
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="revenue.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between rev-target">
                <div>
                  from total target <span style="font-size:15px; font-weight: bold">Rp. {{revenue.targetStr}}</span>
                </div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(revenue.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>

              <q-separator class="q-my-md" />

              <div class="title-subsummary q-mb-md  text-center">
                Total {{selectedTotalTeleTimestamp}} Target Closing
              </div>
              <div class="q-mb-sm">
                <div class="rev-performance row justify-end">
                  <div>{{closing.performance}} <span style="font-size: 16px;color: grey">/ {{closing.target}} Closings</span></div>
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="closing.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between rev-target">
                <div></div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(closing.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>

              <q-separator class="q-my-md" />

              <div class="title-subsummary q-mb-md  text-center">
                Total {{selectedTotalTeleTimestamp}} Target Call
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
    </div>

    <q-separator class="q-my-lg"/>

    <div class="row q-col-gutter-md admin-dashboard">
      <div class="col-md-12">
        <div class="section">
          <div class="title-section text-center">Telemarketer Performance Log Filter</div>
          <q-card>
            <q-card-section class="row">
              <div class="col-md-4 q-pr-md">
                <div class="field-name q-mb-xs">Timestamp</div>
                <q-select
                  outlined
                  v-model="filterLogTimestamp"
                  :options="filterLogTimestampOptions"
                  label="Please select"
                  dense
                  @input="inputLogTimestamp"
                >
                  <template v-if="filterLogTimestamp" v-slot:append>
                    <q-icon name="cancel" @click.stop="clearFilterLogTimestamp" class="cursor-pointer" />
                  </template>
                </q-select>
              </div>
              <div class="col-md-8 q-pl-md">
                <div class="field-name q-mb-xs">Custom Timestamp</div>
                <div class="row no-wrap items-center">
                  <q-input
                    outlined
                    v-model="filterLogCustomTimestampStart"
                    label="Start Date"
                    dense
                    class="full-width"
                    @keydown.enter.prevent="inputLogCustomTimestampStart"
                  >
                    <template v-slot:append>
                      <q-icon name="event" class="cursor-pointer">
                        <q-popup-proxy transition-show="scale" transition-hide="scale">
                          <q-date v-model="filterLogCustomTimestampStart" mask="DD-MMM-YYYY" @input="inputLogCustomTimestampStart"/>
                        </q-popup-proxy>
                      </q-icon>
                    </template>
                  </q-input>
                  <div class="q-mx-sm">to</div>
                  <q-input
                    outlined
                    v-model="filterLogCustomTimestampEnd"
                    label="End Date"
                    dense
                    class="full-width"
                    @keydown.enter.prevent="inputLogCustomTimestampEnd"
                  >
                    <template v-slot:append>
                      <q-icon name="event" class="cursor-pointer">
                        <q-popup-proxy transition-show="scale" transition-hide="scale">
                          <q-date v-model="filterLogCustomTimestampEnd" mask="DD-MMM-YYYY" @input="inputLogCustomTimestampEnd"/>
                        </q-popup-proxy>
                      </q-icon>
                    </template>
                  </q-input>
                </div>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </div>
    <div class="row q-col-gutter-md admin-dashboard">
      <div class="col-md-8">
        <div class="section">
          <div class="title-section text-center">Telemarketer Performance Log</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary q-mb-sm">
                  {{selectedTeleLogType}} {{getLogTimestamp()}}
                </div>
                <div>
                  <q-select outlined v-model="selectedTeleLogType" :options="teleLogTypeOptions" dense style="min-width: 120px" @input="processingTeleLogData"/>
                </div>
              </div>
              <div class="telemarketer-chart">
                <TelemarketerChart v-bind:teleChartData="teleLogChartData" v-if="teleLogChartData.label.length > 0"/>
                <q-spinner-pie
                  class="loading"
                  color="primary"
                  size="50px"
                  v-if="telemarketerLogChartLoading"
                />
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
      <div class="col-md-4">
        <div class="section">
          <div class="title-section text-center">Summary</div>
          <q-card>
            <q-card-section>
              <div class="row justify-between items-center q-mb-md">
                <div class="title-summary">
                  Performance {{getLogTimestamp()}}
                </div>
              </div>
              <div class="title-subsummary q-mb-sm text-center">
                Total Target Revenue
              </div>
              <div class="q-mb-md">
                <div class="rev-performance">
                  <span style="font-size: 24px;">Rp.</span> {{revenueLog.performanceStr}}
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="revenueLog.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between rev-target">
                <div>
                  from total target <span style="font-size:15px; font-weight: bold">Rp. {{revenueLog.targetStr}}</span>
                </div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(revenueLog.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>

              <q-separator class="q-my-md" />

              <div class="title-subsummary q-mb-md  text-center">
                Total Target Closing
              </div>
              <div class="q-mb-sm">
                <div class="rev-performance row justify-end">
                  <div>{{closingLog.performance}} <span style="font-size: 16px;color: grey">/ {{closingLog.target}} Closings</span></div>
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="closingLog.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between rev-target">
                <div></div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(closingLog.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>

              <q-separator class="q-my-md" />

              <div class="title-subsummary q-mb-md  text-center">
                Total Target Call
              </div>
              <div class="q-mb-sm">
                <div class="rev-performance row justify-end">
                  <div>{{callLog.performance}} <span style="font-size: 16px;color: grey">/ {{callLog.target}} Calls</span></div>
                </div>
                <div>
                  <q-linear-progress stripe rounded size="20px" :value="callLog.progress" color="green" class="q-mt-sm" />
                </div>
              </div>
              <div class="row justify-between q-mb-md rev-target">
                <div></div>
                <div>
                  <span style="font-size:15px; font-weight: bold">{{(callLog.progress * 100).toFixed(0)}} %</span>
                </div>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </div>
  </q-page>
</template>

<script>
import { date } from "quasar";
import moment from "moment-timezone";
import TelemarketerChart from "components/TelemarketerChart.vue";

export default {
  name: 'DashboardAdmin',

  components: {
    TelemarketerChart,
  },

  data() {
    return {
      telemarketers: {},
      selectedTeleTimestamp: "Monthly",
      teleTimestampOptions: ["Daily", "Weekly", "Monthly"],
      selectedTeleType: "Revenue",
      teleTypeOptions: ["Revenue", "Closing", "Call"],
      teleChartData: {
        label: [],
        targetColor: [],
        achievementColor: [],
        target: [],
        achievement: [], 
        type: "Revenue",
      },
      selectedTotalTeleTimestamp: "Monthly",
      totalTeleTimestampOptions: ["Daily", "Weekly", "Monthly"],
      revenue: {
        performanceStr: "0",
        targetStr: "0",
        performance: 0,
        target: 0,
        progress: 0,
      },
      closing: {
        performance: 0,
        target: 0,
        progress: 0,
      },
      call: {
        performance: 0,
        target: 0,
        progress: 0,
      },
      telemarketerChartLoading: true,
      filterCustomTimestamp: "",

      // Telemarketer Performance Log
      telemarketersLog: {},
      selectedTeleLogType: "Revenue",
      teleLogTypeOptions: ["Revenue", "Closing", "Call"],
      teleLogChartData: {
        label: [],
        targetColor: [],
        achievementColor: [],
        target: [],
        achievement: [], 
        type: "Revenue",
      },
      revenueLog: {
        performanceStr: "0",
        targetStr: "0",
        performance: 0,
        target: 0,
        progress: 0,
      },
      closingLog: {
        performance: 0,
        target: 0,
        progress: 0,
      },
      callLog: {
        performance: 0,
        target: 0,
        progress: 0,
      },
      telemarketerLogChartLoading: true,
      filterLogTimestamp: "",
      filterLogTimestampOptions: [
        "Today",
        "Yesterday",
        "Last 7 days",
        "Last 15 days",
        "Last 30 days",
        "Last 60 days",
        "Last 90 days",
      ],
      filterLogCustomTimestampStart: "",
      filterLogCustomTimestampEnd: "",
    }
  },

  mounted() {
    this.filterCustomTimestamp = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
    this.filter();
    // Telemarketer Performance Log
    this.filterLogCustomTimestampStart = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
    this.filterLogCustomTimestampEnd = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
    this.filterLog(this.filterLogCustomTimestampStart, this.filterLogCustomTimestampEnd);
  },

  methods: {
    getLogTimestamp(){
      if (this.filterLogTimestamp) {
        if (this.filterLogCustomTimestampStart == this.filterLogCustomTimestampEnd){
          return ": " + this.filterLogTimestamp + " (" + this.filterLogCustomTimestampStart + ")"
        } else {
          return ": " + this.filterLogTimestamp + " (from " + this.filterLogCustomTimestampStart + " to " + this.filterLogCustomTimestampEnd + ")"
        }
      } else {
        if (this.filterLogCustomTimestampStart == this.filterLogCustomTimestampEnd){
          return "in " + this.filterLogCustomTimestampStart
        } else {
          return "from " + this.filterLogCustomTimestampStart + " to " + this.filterLogCustomTimestampEnd
        }
      }
    },
    processingTotalTeleData(){
      var vm = this
      var revenuePerformance = 0
      var revenueTarget = 0
      var closingPerformance = 0
      var closingTarget = 0
      var callPerformance = 0
      var callTarget = 0
      this.telemarketers.forEach(function(data) {
        if (vm.selectedTotalTeleTimestamp == 'Daily'){
          revenuePerformance = revenuePerformance + data.Performance.Daily.BuyAmount
          revenueTarget = revenueTarget + data.Target.Daily.BuyAmount
          closingPerformance = closingPerformance + data.Performance.Daily.Closing
          closingTarget = closingTarget + data.Target.Daily.Closing
          callPerformance = callPerformance + data.Performance.Daily.Call
          callTarget = callTarget + data.Target.Daily.Call
        } else if (vm.selectedTotalTeleTimestamp == 'Weekly') {
          revenuePerformance = revenuePerformance + data.Performance.Weekly.BuyAmount
          revenueTarget = revenueTarget + data.Target.Weekly.BuyAmount
          closingPerformance = closingPerformance + data.Performance.Weekly.Closing
          closingTarget = closingTarget + data.Target.Weekly.Closing
          callPerformance = callPerformance + data.Performance.Weekly.Call
          callTarget = callTarget + data.Target.Weekly.Call
        } else if (vm.selectedTotalTeleTimestamp == 'Monthly') {
          revenuePerformance = revenuePerformance + data.Performance.Monthly.BuyAmount
          revenueTarget = revenueTarget + data.Target.Monthly.BuyAmount
          closingPerformance = closingPerformance + data.Performance.Monthly.Closing
          closingTarget = closingTarget + data.Target.Monthly.Closing
          callPerformance = callPerformance + data.Performance.Monthly.Call
          callTarget = callTarget + data.Target.Monthly.Call
        }
      });
      this.revenue.performanceStr = revenuePerformance.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
      this.revenue.targetStr = revenueTarget.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
      this.revenue.performance = revenuePerformance
      this.revenue.target = revenueTarget
      this.revenue.progress =  revenuePerformance/revenueTarget
      if (revenueTarget == 0) {
        this.revenue.progress = 0
      }
      this.closing.performance = closingPerformance
      this.closing.target = closingTarget
      this.closing.progress =  closingPerformance/closingTarget
      if (closingTarget == 0) {
        this.closing.progress = 0
      }
      this.call.performance = callPerformance
      this.call.target = callTarget
      this.call.progress =  callPerformance/callTarget
      if (callTarget == 0) {
        this.call.progress = 0
      }
    },
    processingTeleData() {
      var vm = this
      var label = []
      var targetColor = []
      var target = []
      var achievementColor = []
      var achievement = []
      this.telemarketers.forEach(function(data) {
        label.push(data.Name)
        targetColor.push('rgba(0, 0, 0, 0.2)')
        achievementColor.push('#2e7d32')
        if (vm.selectedTeleTimestamp == 'Daily') {
          if (vm.selectedTeleType == 'Revenue') {
            target.push(data.Target.Daily.BuyAmount)
            achievement.push(data.Performance.Daily.BuyAmount)
          } else if (vm.selectedTeleType == 'Closing') {
            target.push(data.Target.Daily.Closing)
            achievement.push(data.Performance.Daily.Closing)
          } else if (vm.selectedTeleType == 'Call') {
            target.push(data.Target.Daily.Call)
            achievement.push(data.Performance.Daily.Call)
          }
        } else if (vm.selectedTeleTimestamp == 'Weekly') {
          if (vm.selectedTeleType == 'Revenue') {
            target.push(data.Target.Weekly.BuyAmount)
            achievement.push(data.Performance.Weekly.BuyAmount)
          } else if (vm.selectedTeleType == 'Closing') {
            target.push(data.Target.Weekly.Closing)
            achievement.push(data.Performance.Weekly.Closing)
          } else if (vm.selectedTeleType == 'Call') {
            target.push(data.Target.Weekly.Call)
            achievement.push(data.Performance.Weekly.Call)
          }
        } else if (vm.selectedTeleTimestamp == 'Monthly') {
          if (vm.selectedTeleType == 'Revenue') {
            target.push(data.Target.Monthly.BuyAmount)
            achievement.push(data.Performance.Monthly.BuyAmount)
          } else if (vm.selectedTeleType == 'Closing') {
            target.push(data.Target.Monthly.Closing)
            achievement.push(data.Performance.Monthly.Closing)
          } else if (vm.selectedTeleType == 'Call') {
            target.push(data.Target.Monthly.Call)
            achievement.push(data.Performance.Monthly.Call)
          }
        }
      });
      this.teleChartData = Object.assign({}, this.teleChartData, { 
        label: label,
        targetColor: targetColor,
        achievementColor: achievementColor,
        target: target,
        achievement: achievement, 
        type: vm.selectedTeleType,
      })
      this.telemarketerChartLoading = false
    },
    removeAdminTele(telemarketers) {
      var vm = this
      this.telemarketers = telemarketers.filter(x => x.IsAdmin == false);
      this.processingTotalTeleData()
      setTimeout(function(){
        vm.processingTeleData();
      }, 1000)
    },
    filter() {
      var vm = this
      vm.telemarketers = []
      var targetDate
      targetDate = moment.tz(this.filterCustomTimestamp, "DD-MMM-YYYY HH:mm", "Asia/Jakarta").startOf("day");
      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterTelemarketer: {
          ReportTimestamp: targetDate.unix() * 1000000000,
        },
        Limit: 10000,
      }
      this.$axios
        .post("/api/telemarketer/get", data_submit)
        .then(function (response) {
          if (response.data.Telemarketers != null) {
            vm.$nextTick(() => { 
              vm.removeAdminTele(response.data.Telemarketers)
            })
          }
        })
    },
    inputFilterCustomTimestamp() {
      if(this.filterCustomTimestamp == ""){
        this.filterCustomTimestamp = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
      }
      this.filter();
    },
    
    // Telemarketing Performance Log
    filterLog(from, to){
      var vm = this
      var fromDateTime = moment.tz(from, "DD-MMM-YYYY HH:mm", "Asia/Jakarta").startOf("day");
      var toDateTime = moment.tz(to, "DD-MMM-YYYY HH:mm", "Asia/Jakarta").endOf("day");

      console.log(fromDateTime)
      console.log(toDateTime)

      var data_submit = {
        Token: vm.$authService.getToken(),
        FilterReportTelemarketer: {
          TimestampStart: fromDateTime.unix() * 1000000000,
          TimestampEnd: toDateTime.unix() * 1000000000
        },
        Limit: 10000,
      }
      this.$axios
        .post("/api/telemarketer/report", data_submit)
        .then(function (response) {
          if (response.data.Telemarketers != null) {
            vm.$nextTick(() => { 
              vm.removeAdminTeleLog(response.data.Telemarketers)
            })
          }
        })
      // ALTERNATIVE
      // this.getLogTelemarketerFromAPI(fromDateTime, toDateTime)
    },
    // async getLogTelemarketerFromAPI(fromDateTime, toDateTime){
    //   var vm = this
    //   var teles = []
    //   console.log('Start')
    //   for (var i = fromDateTime.unix(); i < toDateTime.unix(); i += 86400) {
    //     var data_submit = {
    //       Token: vm.$authService.getToken(),
    //       FilterTelemarketer: {
    //         ReportTimestamp: i * 1000000000,
    //       },
    //       Limit: 10000,
    //     }
    //     await this.$axios
    //       .post("/api/telemarketer/get", data_submit)
    //       .then(function (response) {
    //         if (response.data.Telemarketers != null) {
    //           teles.push(response.data.Telemarketers)
    //           console.log(i)
    //           console.log(response.data.Telemarketers)
    //         }
    //       })
    //   }
    //   console.log(teles)
    //   console.log('End')
    // },
    removeAdminTeleLog(telemarketers) {
      var vm = this
      this.telemarketersLog = telemarketers.filter(x => x.IsAdmin == false);
      this.processingTotalTeleLogData()
      setTimeout(function(){
        vm.processingTeleLogData();
      }, 1000)
    },
    processingTotalTeleLogData(){
      var vm = this
      var revenuePerformance = 0
      var revenueTarget = 0
      var closingPerformance = 0
      var closingTarget = 0
      var callPerformance = 0
      var callTarget = 0
      this.telemarketersLog.forEach(function(data) {
        revenuePerformance = revenuePerformance + data.Performance.Daily.BuyAmount
        revenueTarget = revenueTarget + data.Target.Daily.BuyAmount
        closingPerformance = closingPerformance + data.Performance.Daily.Closing
        closingTarget = closingTarget + data.Target.Daily.Closing
        callPerformance = callPerformance + data.Performance.Daily.Call
        callTarget = callTarget + data.Target.Daily.Call
      });
      this.revenueLog.performanceStr = revenuePerformance.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
      this.revenueLog.targetStr = revenueTarget.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
      this.revenueLog.performance = revenuePerformance
      this.revenueLog.target = revenueTarget
      this.revenueLog.progress =  revenuePerformance/revenueTarget
      if (revenueTarget == 0) {
        this.revenueLog.progress = 0
      }
      this.closingLog.performance = closingPerformance
      this.closingLog.target = closingTarget
      this.closingLog.progress =  closingPerformance/closingTarget
      if (closingTarget == 0) {
        this.closingLog.progress = 0
      }
      this.callLog.performance = callPerformance
      this.callLog.target = callTarget
      this.callLog.progress =  callPerformance/callTarget
      if (callTarget == 0) {
        this.callLog.progress = 0
      }
    },
    processingTeleLogData() {
      var vm = this
      var label = []
      var targetColor = []
      var target = []
      var achievementColor = []
      var achievement = []
      this.telemarketersLog.forEach(function(data) {
        label.push(data.Name)
        targetColor.push('rgba(0, 0, 0, 0.2)')
        achievementColor.push('#2e7d32')
        if (vm.selectedTeleLogType == 'Revenue') {
          target.push(data.Target.Daily.BuyAmount)
          achievement.push(data.Performance.Daily.BuyAmount)
        } else if (vm.selectedTeleLogType == 'Closing') {
          target.push(data.Target.Daily.Closing)
          achievement.push(data.Performance.Daily.Closing)
        } else if (vm.selectedTeleLogType == 'Call') {
          target.push(data.Target.Daily.Call)
          achievement.push(data.Performance.Daily.Call)
        }
      });
      this.teleLogChartData = Object.assign({}, this.teleLogChartData, { 
        label: label,
        targetColor: targetColor,
        achievementColor: achievementColor,
        target: target,
        achievement: achievement, 
        type: vm.selectedTeleLogType,
      })
      this.telemarketerLogChartLoading = false
    },
    getFromToTimestamp(filterLogTimestamp) {
      var StartEndDate = {
        from: "",
        to: ""
      }
      if (filterLogTimestamp == "Today") {
        StartEndDate.from = moment().startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (filterLogTimestamp == "Yesterday") {
        StartEndDate.from = moment().subtract(1, "days").startOf("day");
        StartEndDate.to = moment().subtract(1, "days").endOf("day");
      } else if (filterLogTimestamp == "Last 7 days") {
        StartEndDate.from = moment().subtract(7, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (filterLogTimestamp == "Last 15 days") {
        StartEndDate.from = moment().subtract(15, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (filterLogTimestamp == "Last 30 days") {
        StartEndDate.from = moment().subtract(30, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (filterLogTimestamp == "Last 60 days") {
        StartEndDate.from = moment().subtract(60, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      } else if (filterLogTimestamp == "Last 90 days") {
        StartEndDate.from = moment().subtract(90, "days").startOf("day");
        StartEndDate.to = moment().endOf("day");
      }
      return StartEndDate
    },
    inputLogTimestamp() {
      var StartEndDate = this.getFromToTimestamp(this.filterLogTimestamp)
      var fromDate = StartEndDate.from
      var toDate = StartEndDate.to
      this.filterLogCustomTimestampStart = fromDate.format("DD-MMM-YYYY")
      this.filterLogCustomTimestampEnd = toDate.format("DD-MMM-YYYY")
      this.filterLog(this.filterLogCustomTimestampStart, this.filterLogCustomTimestampEnd);
    },
    inputLogCustomTimestampStart() {
      this.filterLogTimestamp = "";
      if(this.filterLogCustomTimestampStart == ""){
        this.filterLogCustomTimestampStart = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
      }
      this.filterLog(this.filterLogCustomTimestampStart, this.filterLogCustomTimestampEnd);
    },
    inputLogCustomTimestampEnd() {
      this.filterLogTimestamp = "";
      if(this.filterLogCustomTimestampEnd == ""){
        this.filterLogCustomTimestampEnd = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
      }
      this.filterLog(this.filterLogCustomTimestampStart, this.filterLogCustomTimestampEnd);
    },
    clearFilterLogTimestamp(){
      this.filterLogTimestamp = "";
      this.filterLogCustomTimestampStart = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
      this.filterLogCustomTimestampEnd = moment.tz("Asia/Jakarta").format("DD-MMM-YYYY");
      this.filterLog(this.filterLogCustomTimestampStart, this.filterLogCustomTimestampEnd);
    },
  }
}
</script>
