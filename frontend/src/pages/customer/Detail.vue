<template>
  <q-page class="customer-detail">
    <div class="row items-center q-mb-md">
      <q-btn round color="blue-grey" icon="chevron_left" size="sm" @click="backToIndex" />
      <div text-color="blue-grey" class="q-ml-sm text-bold" style="font-size:16px;">Back</div>
    </div>
    <div class="row">
      <div class="title">Customer | Phone: {{customerPhoneNumber}}</div>
    </div>
    <div class="detail-container">
      <q-form @submit="onSubmit">
        <q-card class="row q-pa-md">
          <!--Customer Information-->
          <q-card class="customer-info col-md-4">
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
                />
              </div>
              <!--<div>
                <div class="field-name q-mb-xs">Additonal Information </div>
                <q-input filled v-model="additionalInfo" type="textarea"/>
              </div>-->
            </q-card-section>
          </q-card>

          <!--Customer Response-->
          <q-card class="customer-response col-md-8">
            <div class="q-pl-lg full-height">
              <div class="column justify-between full-height">
                <div>
                  <div class="q-mb-md">
                    <div class="title">Call by</div>
                    <div>
                      <q-btn
                        style="background-color: #00aff0"
                        text-color="white"
                        icon="fab fa-skype"
                        label="Skype"
                        class="q-px-sm"
                        no-caps
                        @click="callBySkype"
                      />
                    </div>
                  </div>
                  <div>
                    <div class="title">
                      Response
                      <span style="color: red; font-weight: normal">*</span>
                    </div>
                    <div class="row q-col-gutter-sm">
                      <div class="col-md-4">
                        <q-btn
                          outline
                          icon="fas fa-smile"
                          stack
                          color="green-9"
                          label="Tertarik"
                          class="full-width q-py-sm"
                          no-caps
                          v-bind:class="{ 'active-tertarik': tertarik }"
                          @click="activateTertarik"
                        />
                      </div>
                      <div class="col-md-4">
                        <q-btn
                          outline
                          icon="fas fa-reply"
                          stack
                          color="blue-9"
                          label="Hubungi Kembali"
                          class="full-width q-py-sm"
                          no-caps
                          v-bind:class="{ 'active-hubungi-kembali': hubungiKembali }"
                          @click="activateHubungiKembali"
                        />
                      </div>
                      <div class="col-md-4">
                        <q-btn
                          outline
                          icon="fas fa-frown"
                          stack
                          color="red-9"
                          label="Tidak Tertarik"
                          class="full-width q-py-sm"
                          no-caps
                          v-bind:class="{ 'active-tidak-tertarik': tidakTertarik }"
                          @click="activateTidakTertarik"
                        />
                      </div>
                      <div class="col-md-4">
                        <q-btn
                          outline
                          icon="fas fa-times"
                          stack
                          color="grey-6"
                          label="Tidak Aktif"
                          class="full-width q-py-sm"
                          no-caps
                          v-bind:class="{ 'active-tidak-aktif': tidakAktif }"
                          @click="activateTidakAktif"
                        />
                      </div>
                      <div class="col-md-4">
                        <q-btn
                          outline
                          icon="fas fa-share"
                          stack
                          color="orange-7"
                          label="Tidak Menjawab"
                          class="full-width q-py-sm"
                          no-caps
                          v-bind:class="{ 'active-tidak-menjawab': tidakMenjawab }"
                          @click="activateTidakMenjawab"
                        />
                      </div>
                      <div class="col-md-4">
                        <q-btn
                          outline
                          icon="fas fa-ban"
                          stack
                          color="black"
                          label="Tidak Terdaftar"
                          class="full-width q-py-sm"
                          no-caps
                          v-bind:class="{ 'active-tidak-terdaftar': tidakTerdaftar }"
                          @click="activateTidakTerdaftar"
                        />
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-end q-mt-lg">
                  <q-btn
                    class="q-px-lg"
                    color="primary"
                    text-color="white"
                    type="submit"
                    label="Save"
                  />
                </div>
              </div>
            </div>
          </q-card>
        </q-card>
      </q-form>
    </div>
    <div class="table-container">
      <q-table
        :data="callHistory"
        :columns="callHistoryColumns"
        row-key="Timestamp"
        :filter="callHistoryFilter"
        :visible-columns="callHistoryVisible"
        :pagination.sync="callHistoryPagination"
      >
        <template v-slot:top-right>
          <q-input color="grey-8" dense debounce="300" v-model="callHistoryFilter" placeholder="Search">
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>
        <template v-slot:top-left>
          <div class="q-mb-sm">
            <span class="q-pr-sm">Total Call History</span>
            <span class="text-bold" style="font-size: 14px">{{callHistory.length}}</span>
          </div>
        </template>
      </q-table>
    </div>
    <q-dialog v-model="alert">
      <q-card>
        <q-card-section>
          <div class="text-h6 text-center">Alert</div>
        </q-card-section>

        <q-card-section class="q-pt-none">Response is required</q-card-section>
      </q-card>
    </q-dialog>
  </q-page>
</template>

<script>
import { date } from "quasar";

export default {
  name: "CustomerDetail",

  data() {
    return {
      customerPhoneNumber: this.$route.params.phoneNumber,
      customer: {},
      additionalInfo: "",
      response: "",
      alert: false,

      tertarik: false,
      hubungiKembali: false,
      tidakTertarik: false,
      tidakAktif: false,
      tidakMenjawab: false,
      tidakTerdaftar: false,

      callHistory: [],
      callHistoryColumns: [
        {
          name: "date",
          label: "DATE",
          align: "left",
          field: "Date",
          sortable: true,
        },
        {
          name: "name",
          label: "NAME",
          align: "left",
          field: "Name",
          sortable: true,
        },
        {
          name: "phoneNumber",
          label: "PHONE NUMBER",
          align: "center",
          field: "PhoneNumber",
          sortable: true,
        },
        {
          name: "status",
          label: "STATUS",
          align: "center",
          field: "Status",
          sortable: true,
        },
      ],
      callHistoryFilter: "",
      callHistoryVisible: ["date", "name", "phoneNumber", "status"],
      callHistoryPagination: {
        rowsPerPage: 5, // current rows per page being displayed
      },
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
    // Assign Status if there's a value in it
    this.assignStatus(this.customer.Status);

    var data_submit = {
      Token: vm.$authService.getToken(),
      FilterCallLog: {
        PhoneNumber: vm.customerPhoneNumber,
      },
      Limit: 10000,
    };
    this.$axios
      .post("/api/call-log/get", data_submit)
      .then(function (response) {
        if (response.data) {
          if (response.data.CallLogs) {
            vm.assignDataFromAPI(response.data.CallLogs);
          } else {
            vm.callHistory = [];
          }
        }
      });
  },

  methods: {
    assignDataFromAPI(dataResponse) {
      var vm = this;
      var data = dataResponse;
      Object.keys(data).forEach((key) => {
        var Timestamp = new Date(data[key].Timestamp / 1000000);
        Timestamp = date.formatDate(Timestamp, "DD-MM-YYYY HH:mm:ss Z");

        var dataModel = data[key];
        dataModel.Date = Timestamp;

        vm.$set(vm.callHistory, key, dataModel);
      });
    },
    backToIndex() {
      this.$router.go(-1);
    },
    callBySkype() {
      // console.log("Call by Skype")
      window.location = "skype://+" + this.customer.PhoneNumber + "?call";
    },
    onSubmit() {
      var vm = this;
      // console.log(this.customer);
      // console.log(this.additionalInfo);
      // console.log(this.response);
      if (this.response == "") {
        this.alert = true;
      } else {
        // Submit data
        var data_submit = {
          Token: vm.$authService.getToken(),
          PhoneNumber: vm.customerPhoneNumber,
          Status: vm.response,
        };
        vm.$restapi.callCustomer(data_submit, ()=>{vm.$router.push({ name: "customer" })});
      }
    },
    assignStatus(status) {
      if (status == "Tertarik") {
        this.tertarik = true;
        this.response = "Tertarik";
      } else if (status == "Hubungi Kembali") {
        this.hubungiKembali = true;
        this.response = "Hubungi Kembali";
      } else if (status == "Tidak Tertarik") {
        this.tidakTertarik = true;
        this.response = "Tidak Tertarik";
      } else if (status == "Tidak Aktif") {
        this.tidakAktif = true;
        this.response = "Tidak Aktif";
      } else if (status == "Tidak Menjawab") {
        this.tidakMenjawab = true;
        this.response = "Tidak Menjawab";
      } else if (status == "Tidak Terdaftar") {
        this.tidakTerdaftar = true;
        this.response = "Tidak Terdaftar";
      }
    },
    reset() {
      this.tertarik = false;
      this.hubungiKembali = false;
      this.tidakTertarik = false;
      this.tidakAktif = false;
      this.tidakMenjawab = false;
      this.tidakTerdaftar = false;
    },
    activateTertarik() {
      if (this.tertarik) {
        this.tertarik = false;
        this.response = "";
      } else {
        this.reset();
        this.tertarik = true;
        this.response = "Tertarik";
      }
    },
    activateHubungiKembali() {
      if (this.hubungiKembali) {
        this.hubungiKembali = false;
        this.response = "";
      } else {
        this.reset();
        this.hubungiKembali = true;
        this.response = "Hubungi Kembali";
      }
    },
    activateTidakTertarik() {
      if (this.tidakTertarik) {
        this.tidakTertarik = false;
        this.response = "";
      } else {
        this.reset();
        this.tidakTertarik = true;
        this.response = "Tidak Tertarik";
      }
    },
    activateTidakAktif() {
      if (this.tidakAktif) {
        this.tidakAktif = false;
        this.response = "";
      } else {
        this.reset();
        this.tidakAktif = true;
        this.response = "Tidak Aktif";
      }
    },
    activateTidakMenjawab() {
      if (this.tidakMenjawab) {
        this.tidakMenjawab = false;
        this.response = "";
      } else {
        this.reset();
        this.tidakMenjawab = true;
        this.response = "Tidak Menjawab";
      }
    },
    activateTidakTerdaftar() {
      if (this.tidakTerdaftar) {
        this.tidakTerdaftar = false;
        this.response = "";
      } else {
        this.reset();
        this.tidakTerdaftar = true;
        this.response = "Tidak Terdaftar";
      }
    },
  },
};
</script>
