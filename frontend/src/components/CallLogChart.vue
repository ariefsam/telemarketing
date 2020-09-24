<script>
  import { Doughnut } from 'vue-chartjs'

  export default {
    name: 'CallLogChart',

    extends: Doughnut,

    data () {
      return {
        chartData: {},
        options: {
          legend: {
            display: false
          },
          responsive: true,
          maintainAspectRatio: false,
        },
      }
    },

    mounted () {
      this.generateEmptyChart()    
    },

    methods: {
      generateEmptyChart() {
        this.chartData = {
          labels: ["Tertarik",	"Hubungi Kembali",	"Tidak Tertarik",	"Tidak Aktif", "Tidak Menjawab", "Tidak Terdaftar", "HOT 80%"],
          datasets: [
            {
              backgroundColor: [
                'rgba(0, 0, 0, 0.05)',      
                'rgba(0, 0, 0, 0.05)',      
                'rgba(0, 0, 0, 0.05)',      
                'rgba(0, 0, 0, 0.05)',      
                'rgba(0, 0, 0, 0.05)',      
                'rgba(0, 0, 0, 0.05)',      
                'rgba(0, 0, 0, 0.05)'   
              ],
              data: [1,1,1,1,1,1,1],
            }
          ]
        }
        this.options.tooltips = {
          enabled: false
        }
        this.renderChart(this.chartData, this.options)
      },
      generateCallLogChart(callLogStatus) {
        this.chartData = {
          labels: ["Tertarik",	"Hubungi Kembali",	"Tidak Tertarik",	"Tidak Aktif", "Tidak Menjawab", "Tidak Terdaftar", "HOT 80%"],
          datasets: [
            {
              backgroundColor: [
              '#2e7d32',
              '#1565c0',
              '#c62828',
              '#9e9e9e',
              '#fb8c00',                
              '#000000',   
              '#673ab7'     
              ],
              data: [
                callLogStatus.tertarik, 
                callLogStatus.hubungiKembali, 
                callLogStatus.tidakTertarik,
                callLogStatus.tidakAktif,
                callLogStatus.tidakMenjawab,
                callLogStatus.tidakTerdaftar,
                callLogStatus.hot
              ],
            }
          ]
        }
        this.options.tooltips = {
          enabled: true
        }
        this.renderChart(this.chartData, this.options)
      }
    },
    
    watch: {
      callLogStatus: {
        handler(newVal){
          if (newVal.tertarik == 0 && newVal.hubungiKembali == 0 && newVal.tidakTertarik == 0 && newVal.tidakAktif == 0 && newVal.tidakMenjawab == 0 && newVal.tidakTerdaftar == 0 && newVal.hot == 0) {
            this.generateEmptyChart() 
          } else {
            this.generateCallLogChart(newVal)
          }
        },
        deep: true
      }
    },

    props: {
      callLogStatus: {
        type: Object,
        required: true
      },
    }
  }
</script>