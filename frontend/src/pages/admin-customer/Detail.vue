<template>
  <q-page class="customer-detail">
    <div class="row items-center q-mb-md">
      <q-btn round color="blue-grey" icon="chevron_left" size="sm" @click="backToIndex" />
      <div text-color="blue-grey" class="q-ml-sm text-bold" style="font-size:16px;">Back</div>
    </div>
    <div class="row">
      <div class="title">Customer | Phone: {{customerPhoneNumber}}</div>
    </div>
    <div class="detail-container row">
      <q-card class="q-pa-md col-md-4">
        <!--Customer Information-->
        <q-card class="customer-info">
          <q-card-section class="title">
            <div class="text-h6">Customer Information</div>
          </q-card-section>
          <q-separator />
          <q-card-section class="content">
            <div class="q-mb-xs">
              <div class="field-name q-mb-xs">
                Name
                <span style="color: red; font-weight: normal">*</span>
              </div>
              <q-input
                filled
                v-model="customer.Name"
                dense
                :rules="[val => !!val || 'Field is required']"
                readonly
              />
            </div>
            <div class="q-mb-xs">
              <div class="field-name q-mb-xs">
                Phone Number
                <span style="color: red; font-weight: normal">*</span>
              </div>
              <q-input
                filled
                v-model="customer.PhoneNumber"
                dense
                :rules="[val => !!val || 'Field is required']"
                readonly
              />
            </div>
            <!--<div>
              <div class="field-name q-mb-xs">Additonal Information </div>
              <q-input filled v-model="additionalInfo" type="textarea"/>
            </div>-->
          </q-card-section>
        </q-card>
      </q-card>
    </div>
  </q-page>
</template>


<script>
export default {
  name: "AdminCustomerDetail",

  data() {
    return {
      customerPhoneNumber: this.$route.params.phoneNumber,
      customer: {},
      additionalInfo: "",
    };
  },

  mounted() {
    var vm = this;
    var data_submit = {
      Token: vm.$authService.getToken(),
      FilterCustomer: {
        PhoneNumber: vm.customerPhoneNumber,
      },
      Limit: 10000,
    };
    this.$axios.post("/api/customer", data_submit).then(function (response) {
      if (response.data) {
        vm.customer = response.data.Customers[0];
      }
    });
  },

  methods: {
    backToIndex() {
      this.$router.replace({ name: "admin-customer" });
    },
  },
};
</script>
