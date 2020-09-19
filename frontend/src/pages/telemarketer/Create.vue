<template>
  <q-page class="page-index">
    <div class="row items-center q-mb-md">
      <q-btn round color="blue-grey" icon="chevron_left" size="sm" @click="backToIndex" />
      <div text-color="blue-grey" class="q-ml-sm text-bold" style="font-size:16px;">Back</div>
    </div>
    <div class="row">
      <div class="title">Add Telemarketer</div>
    </div>
    <div class="row">
      <q-form @submit="onSubmit" class="col-md-8">
        <q-input
          color="grey-8"
          v-model="name"
          filled
          label="Name *"
          type="text"
          class="field"
          :rules="[val => !!val || 'Field is required']"
        />
        <q-input
          color="grey-8"
          v-model="email"
          filled
          label="Email Address *"
          type="email"
          class="field"
          :rules="[val => !!val || 'Field is required']"
        />
        <div class="q-mb-sm">
          <q-toggle v-model="isAdmin" checked-icon="check" color="green" unchecked-icon="clear" />Admin
        </div>
        <div v-if="!isAdmin">
          <div class="q-mb-sm">
            Daily Target
          </div>
          <div class="row q-col-gutter-sm">
            <q-input
              color="grey-8"
              v-model.number="target.Daily.Call"
              filled
              label="Call *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model.number="target.Daily.Closing"
              filled
              label="Closing *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model="target.Daily.BuyAmount"
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
              v-model.number="target.Weekly.Call"
              filled
              label="Call *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model.number="target.Weekly.Closing"
              filled
              label="Closing *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model="target.Weekly.BuyAmount"
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
              v-model.number="target.Monthly.Call"
              filled
              label="Call *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model.number="target.Monthly.Closing"
              filled
              label="Closing *"
              type="number"
              class="field col-md-3"
              :rules="[val => !!val || 'Field is required']"
            />
            <q-input
              color="grey-8"
              v-model="target.Monthly.BuyAmount"
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
  name: "TelemarketerCreate",

  data() {
    return {
      name: "",
      email: "",
      isAdmin: false,
      target: {
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
      },
    };
  },

  watch: {
    target: {
      handler(val){
        const resultDaily = val.Daily.BuyAmount.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultWeekly = val.Weekly.BuyAmount.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        const resultMonthly = val.Monthly.BuyAmount.replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".");
        this.$nextTick(() => { 
          this.target.Daily.BuyAmount = resultDaily 
          this.target.Weekly.BuyAmount = resultWeekly 
          this.target.Monthly.BuyAmount = resultMonthly 
        });
      },
      deep: true
    }
  },

  methods: {
    onSubmit() {
      var vm = this;
      if (vm.isAdmin) {
        vm.target = {
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
      }else {
        vm.target.Daily.BuyAmount = parseInt(vm.target.Daily.BuyAmount.replace(/\./g,''))
        vm.target.Weekly.BuyAmount = parseInt(vm.target.Weekly.BuyAmount.replace(/\./g,''))
        vm.target.Monthly.BuyAmount = parseInt(vm.target.Monthly.BuyAmount.replace(/\./g,''))
      }
      
      var data_submit = {
        Token: vm.$authService.getToken(),
        Telemarketer: {
          Name: vm.name,
          Email: vm.email,
          IsAdmin: vm.isAdmin,
          Target: vm.target
        },
      };
      this.$axios
        .post("/api/telemarketer/create", data_submit)
        .then(function (response) {
          if (response.data) {
            vm.$router.replace({
              name: "telemarketer",
              force: true
            });
          }
        })
        .catch(function (error) {
          vm.$q.dialog({
            title: "Error Create Telemarketer",
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