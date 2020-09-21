<script>
  import { Bar } from 'vue-chartjs'

  export default {
    name: 'TelemarketerChart',

    extends: Bar,

    data() {
      return {
        chartData: {},
        options: {
          scales: {
            yAxes: [
              {
                ticks: {
                  beginAtZero: true
                },
                gridLines: {
                  display: true
                }
              }
            ],
            xAxes: [
              {
                stacked: true,
                id: "target",
                gridLines: {
                  display: false
                },
              },
              {
                display: false,
                stacked: true,
                id: "achievement",
                gridLines: {
                  display: false
                },
                offset: true
              }
            ]
          },
          legend: {
            display: true
          },
          responsive: true,
          maintainAspectRatio: false,
        }
      }
    },

    mounted() {
      this.generateChart(this.teleChartData)
    },

    methods: {
      generateChart(teleChartData) {
        this.chartData = {
          labels: teleChartData.label,
          datasets: [
            {
              label: 'Achievement',
              backgroundColor: teleChartData.achievementColor,
              data: teleChartData.achievement,
              xAxisID: "achievement",
              barThickness: 15,
            },
            {
              label: 'Target',
              backgroundColor: teleChartData.targetColor,
              data: teleChartData.target,
              xAxisID: "target",
              barThickness: 30,
            }
          ]
        }
        this.options.tooltips = {
          enabled: true,
          callbacks: {
            label: function(tooltipItem, data) {
              if (teleChartData.type == "Revenue") {
                return "Rp." + tooltipItem.yLabel.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
              } else {
                return tooltipItem.yLabel
              }
            }
          }
        }
        this.options.scales.yAxes[0] = {
          ticks: {
            beginAtZero: true,
            callback: function(label, index, labels) {
              if (teleChartData.type == "Revenue") {
                return "Rp." + label.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
              } else {
                return label
              }
            },
          },
          gridLines: {
            display: true
          }
        }
        this.renderChart(this.chartData, this.options)
      },
    },

    watch: {
      teleChartData: {
        handler(newVal){
          this.generateChart(newVal)
        },
        deep: true
      }
    },

    props: {
      teleChartData: {
        type: Object,
        required: true
      },
    }
  }
</script>