<template>
  <q-page class="page-index">
    <div class="title">Dashboard</div>
    <div class="row q-col-gutter-md telemarketer-dashboard">
      <div class="col-md-7">
        <q-card>
          <q-card-section>
            <div class="row justify-between">
              <div class="title-summary">
                My Target Revenue
              </div>
              <div>
                Button
              </div>
            </div>
            <div>
              <div class="rev-performance">
                Rp.
              </div>
            </div>
          </q-card-section>
        </q-card>
      </div>
      <div class="col-md-5">
        <q-card>
          <q-card-section>
            2
          </q-card-section>
        </q-card>
      </div>
    </div>
  </q-page>
</template>

<script>
export default {
  name: 'DashboardTelemarketer',

  data() {
    return {
      telemarketer: {
        Target: {
          Daily: {
            Call: "",
            Closing: "",
            BuyAmount: "",
          },
          Weekly: {
            Call: "",
            Closing: "",
            BuyAmount: "",
          },
          Monthly: {
            Call: "",
            Closing: "",
            BuyAmount: "",
          },
        }
      },
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
      .post("/api/telemarketer/get", data_submit)
      .then(function (response) {
        if (response.data) {
          vm.telemarketer = response.data.Telemarketers[0]
        }
      })
      .catch(function(error) {
        console.log(error)
      });

  }
}
</script>
