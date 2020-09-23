<template>
  <q-page class="page-index">
    <div class="title">Dashboard</div>
    <div class="row q-col-gutter-md admin-dashboard">
      <div class="col-md-8">
        <div class="section">
          <div class="title-section text-center">Telemarketer Performance</div>
          <q-card>
            <q-card-section>
              <div class="column justify-between items-center q-mb-md">
                <div class="title-summary q-mb-sm">
                  Telemarketer {{selectedTeleTimestamp}} {{selectedTeleType}}
                </div>
                <div class="row justify-between full-width">
                  <div class="row no-wrap">
                    <!--<div class="q-mr-sm">
                      <q-select
                        outlined
                        v-model="filterTimestamp"
                        :options="filterTimestampOptions"
                        label="Timestamp"
                        dense
                        style="min-width: 165px"
                        @input="inputFilterTimestamp"
                      >
                        <template v-if="filterTimestamp" v-slot:append>
                          <q-icon name="cancel" @click.stop="clearFilterTimestamp" class="cursor-pointer" />
                        </template>
                      </q-select>
                    </div>-->
                    <div>
                      <q-input
                        outlined
                        v-model="filterCustomTimestamp"
                        label="Custom Timestamp"
                        dense
                        style="width: 165px"
                        @keydown.enter.prevent="inputFilterCustomTimestamp"
                      >
                        <template v-slot:append>
                          <q-icon name="event" class="cursor-pointer">
                            <q-popup-proxy transition-show="scale" transition-hide="scale">
                              <q-date v-model="filterCustomTimestamp" mask="DD-MM-YYYY" @input="inputFilterCustomTimestamp"/>
                            </q-popup-proxy>
                          </q-icon>
                        </template>
                      </q-input>
                    </div>
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
      colorList: [],

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

      // filterTimestamp: "",
      // filterTimestampOptions: [
      //   "Today",
      //   "Yesterday",
      //   "Last 7 days",
      //   "Last 15 days",
      //   "Last 30 days",
      //   "Last 60 days",
      //   "Last 90 days",
      // ],
      filterCustomTimestamp: "",
    }
  },

  mounted() {
    this.filterCustomTimestamp = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
    this.filter();
    // var vm = this;
    // var data_submit = {
    //   Token: vm.$authService.getToken(),
    //   Limit: 10000,
    // }
    // this.$axios
    //   .post("/api/telemarketer/get", data_submit)
    //   .then(function (response) {
    //     if (response.data) {
    //       vm.$nextTick(() => { 
    //         vm.removeAdminTele(response.data.Telemarketers)
    //       })
    //     }
    //   })
  },

  methods: {
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
      var achievement = []
      this.telemarketers.forEach(function(data) {
        label.push(data.Name)
        targetColor.push('rgba(0, 0, 0, 0.2)')
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
        achievementColor: vm.colorList,
        target: target,
        achievement: achievement, 
        type: vm.selectedTeleType,
      })
      this.telemarketerChartLoading = false
    },
    removeAdminTele(telemarketers) {
      var vm = this
      this.telemarketers = telemarketers.filter(x => x.IsAdmin == false);
      var colorList = []
      this.telemarketers.forEach(function(data) {
        // colorList.push("#"+vm.randomColor())
        colorList.push('#2e7d32')
      });
      this.colorList = colorList
      this.processingTotalTeleData()
      setTimeout(function(){
        vm.processingTeleData();
      }, 1000)
    },
    randomColor() {
      return Math.floor(Math.random() * 16777215).toString(16);
    },

    // getDate(){
    //  console.log("get Date from filterTimestampOptions") 
    // },
    filter() {
      var vm = this
      vm.telemarketers = []
      var targetDate

      // if (this.filterTimestamp) {
      //   console.log("OPTION DATE")
      //   targetDate = this.getDate(this.filterTimestamp)
      // } else {
      //   console.log("CUSTOM DATE")
      //   targetDate = moment.tz(this.filterCustomTimestamp, "DD-MM-YYYY HH:mm", "Asia/Jakarta").startOf("day");
      // }
      // console.log(targetDate.unix() * 1000000000)

      targetDate = moment.tz(this.filterCustomTimestamp, "DD-MM-YYYY HH:mm", "Asia/Jakarta").startOf("day");
      console.log(targetDate)

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
          if (response.data) {
            vm.$nextTick(() => { 
              vm.removeAdminTele(response.data.Telemarketers)
            })
          }
        })
    },
    // inputFilterTimestamp() {
    //   this.filter();
    // },
    // clearFilterTimestamp() {
    //   this.filterTimestamp = "";
    //   this.filterCustomTimestamp = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
    //   this.filter();
    // },
    inputFilterCustomTimestamp() {
      // this.filterTimestamp = "";
      if(this.filterCustomTimestamp == ""){
        this.filterCustomTimestamp = moment.tz("Asia/Jakarta").format("DD-MM-YYYY");
      }
      this.filter();
    },
  }
}
</script>
