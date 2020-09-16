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
        <div class="q-mb-sm" v-if="!telemarketer.IsAdmin">
          Weekly Target
        </div>
        <div class="row q-col-gutter-sm" v-if="!telemarketer.IsAdmin">
          <q-input
            color="grey-8"
            v-model.number="callTarget"
            filled
            label="Call Target *"
            type="number"
            class="field col-md-6"
            :rules="[val => !!val || 'Field is required']"
          />
          <q-input
            color="grey-8"
            v-model.number="closingTarget"
            filled
            label="Closing Target *"
            type="number"
            class="field col-md-6"
            :rules="[val => !!val || 'Field is required']"
          />
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
      telemarketer: {},
      callTarget: null,
      closingTarget: null,
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
          vm.callTarget = vm.telemarketer.WeeklyTarget.Call
          vm.closingTarget = vm.telemarketer.WeeklyTarget.Closing
        }
      })
  },

  methods: {
    onSubmit() {
      var vm = this;
      if (vm.telemarketer.IsAdmin){
        vm.telemarketer.WeeklyTarget.Call = null
        vm.telemarketer.WeeklyTarget.Closing = null
      } else {
        vm.telemarketer.WeeklyTarget.Call = vm.callTarget
        vm.telemarketer.WeeklyTarget.Closing = vm.closingTarget
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