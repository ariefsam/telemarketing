<template>
  <q-page class="page-index">
    <div class="row items-center q-mb-md">
      <q-btn round color="blue-grey" icon="chevron_left" size="sm" @click="backToIndex" />
      <div text-color="blue-grey" class="q-ml-sm text-bold" style="font-size:16px;">Back</div>
    </div>
    <div class="row">
      <div class="title">Edit Telemarketer</div>
    </div>
    <div class="row">
      <q-form @submit="onSubmit" class="col-md-6">
        <q-input
          color="grey-8"
          v-model="telemarketer.Name"
          filled
          label="Name *"
          type="text"
          class="field"
          :rules="[val => !!val || 'Field is required']"
        />
        <q-input
          color="grey-8"
          v-model="telemarketer.Email"
          filled
          label="Email Address *"
          type="email"
          class="field"
          :rules="[val => !!val || 'Field is required']"
        />
        <div class="q-mb-sm">
          <q-toggle v-model="telemarketer.IsAdmin" checked-icon="check" color="green" unchecked-icon="clear" />Admin
        </div>
        <div v-if="!telemarketer.IsAdmin">
          <div class="q-mb-sm">
            Daily Target
          </div>
          <div class="row q-col-gutter-sm">
            <q-input
              color="grey-8"
              v-model.number="telemarketer.Target.Daily.Call"
              filled
              label="Call *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model.number="telemarketer.Target.Daily.Closing"
              filled
              label="Closing *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model="telemarketer.Target.Daily.BuyAmount"
              filled
              label="Deposit Amount (Rp.) *"
              class="field col-md-6"
              :rules="[val => !!val || 'Field is required']"
            />
          </div>
          <div class="q-mb-sm">
            Weekly Target
          </div>
          <div class="row q-col-gutter-sm">
            <q-input
              color="grey-8"
              v-model.number="telemarketer.Target.Weekly.Call"
              filled
              label="Call *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model.number="telemarketer.Target.Weekly.Closing"
              filled
              label="Closing *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model="telemarketer.Target.Weekly.BuyAmount"
              filled
              label="Deposit Amount (Rp.) *"
              class="field col-md-6"
              :rules="[val => !!val || 'Field is required']"
            />
          </div>
          <div class="q-mb-sm">
            Monthly Target
          </div>
          <div class="row q-col-gutter-sm">
            <q-input
              color="grey-8"
              v-model.number="telemarketer.Target.Monthly.Call"
              filled
              label="Call *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model.number="telemarketer.Target.Monthly.Closing"
              filled
              label="Closing *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model="telemarketer.Target.Monthly.BuyAmount"
              filled
              label="Deposit Amount (Rp.) *"
              class="field col-md-6"
              :rules="[val => !!val || 'Field is required']"
            />
          </div>
        </div>
        <div class="q-mt-sm">
          <q-btn label="Submit" type="submit" color="primary" />
        </div>
      </q-form>
    </div>
  </q-page>
</template>

<script>
import { Dialog } from "quasar";

export default {
  name: "TelemarketerEdit",

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
    };
  },

  mounted(){
    var vm=this;
    var data_submit = {
      Token: vm.$authService.getToken(),
      FilterTelemarketer: {
        ID: vm.$route.params.id,
      },
      Limit: 10000,
    };
    this.$axios
      .post("/api/telemarketer/get", data_submit)
      .then(function (response) {
        if (response.data) {
          vm.telemarketer = response.data.Telemarketers[0]
        }
      })
  },

  watch: {
    telemarketer: {
      handler(val){
        var vm = this
        const resultDaily = val.Target.Daily.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultWeekly = val.Target.Weekly.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultMonthly = val.Target.Monthly.BuyAmount.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        this.$nextTick(() => { 
          vm.telemarketer.Target.Daily.BuyAmount = resultDaily 
          vm.telemarketer.Target.Weekly.BuyAmount = resultWeekly 
          vm.telemarketer.Target.Monthly.BuyAmount = resultMonthly 
        });
      },
      deep: true
    }
  },

  methods: {
    onSubmit() {
      var vm = this;
      if (vm.telemarketer.IsAdmin) {
        vm.telemarketer.Target = {
          Daily: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
          },
          Weekly: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
          },
          Monthly: {
            Call: 0,
            Closing: 0,
            BuyAmount: 0,
          },
        }
      } else {
        vm.telemarketer.Target.Daily.BuyAmount = parseInt(vm.telemarketer.Target.Daily.BuyAmount.replace(/\./g,''))
        vm.telemarketer.Target.Weekly.BuyAmount = parseInt(vm.telemarketer.Target.Weekly.BuyAmount.replace(/\./g,''))
        vm.telemarketer.Target.Monthly.BuyAmount = parseInt(vm.telemarketer.Target.Monthly.BuyAmount.replace(/\./g,''))
      }

      var data_submit = {
        Token: vm.$authService.getToken(),
        Telemarketer: vm.telemarketer,
      };
      this.$axios
        .post("/api/telemarketer/save", data_submit)
        .then(function (response) {
          if (response.data) {
            vm.$router.replace({
              name: "telemarketer",
            });
          }
        })
        .catch(function (error) {
          vm.$q.dialog({
            title: "Error Edit Telemarketer",
            message: error.response.data.Error,
            cancel: true,
            persistent: true,
          });
        });
      console.log(data_submit);
    },
    backToIndex() {
      this.$router.replace({ name: "telemarketer" });
    },
  },
};
</script>