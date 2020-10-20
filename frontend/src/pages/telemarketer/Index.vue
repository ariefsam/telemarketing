<template>
  <q-page class="page-index">
    <div class="row justify-between">
      <div class="title">Telemarketer</div>
      <div>
        <q-btn
          color="light-green-9"
          text-color="white"
          icon="add"
          label="Add"
          class="q-pr-sm"
          :to="{ name: 'telemarketer-create' }"
        />
      </div>
    </div>
    <div class="table-container">
      <q-table
        :data="telemarketers"
        :columns="telemarketerDataColumns"
        row-key="email"
        :filter="telemarketerDataFilter"
        :visible-columns="telemarketerDataVisible"
        :pagination.sync="telemarketerDataPagination"
        @row-click="onRowClick"
      >
        <template v-slot:top-right>
          <div class="column">
            <q-input color="grey-8" dense debounce="300" v-model="telemarketerDataFilter" placeholder="Search" class="q-mb-sm">
              <template v-slot:append>
                <q-icon name="search" />
              </template>
            </q-input>
            <q-select
              filled
              v-model="telemarketerDataVisible"
              multiple
              outlined
              dense
              :options="telemarketerDataColumns"
              :display-value="$q.lang.table.columns"
              emit-value
              map-options
              option-value="name"
              style="min-width: 240px"
            />
          </div>
        </template>
        <template v-slot:top-left>
          <div>
            <span class="q-pr-sm">Total Telemarketer</span>
            <span class="text-bold" style="font-size: 14px">{{telemarketers.length}}</span>
          </div>
        </template>
        <q-td slot="body-cell-isAdmin" slot-scope="props" :props="props">
          <div v-if="props.value == true">Yes</div>
          <div v-else>No</div>
        </q-td>
        <q-td slot="body-cell-action" slot-scope="props" :props="props">
          <q-btn
            color="primary"
            icon="edit"
            class="action-btn q-mx-xs"
            @click.stop="editTelemarketer(props.row.ID)"
            padding="sm"
          >
            <q-tooltip>Edit Telemarketer</q-tooltip>
          </q-btn>
          <q-btn
            color="red"
            icon="fas fa-trash"
            class="action-btn q-mx-xs"
            @click.stop="confirmDeleteTelemarketer(props.row)"
            padding="sm"
          >
            <q-tooltip>Delete Telemarketer</q-tooltip>
          </q-btn>
        </q-td>
      </q-table>
    </div>
  </q-page>
</template>

<script>
export default {
  name: 'TelemarketerIndex',

  data() {
    return {
      telemarketers: [],
      telemarketerDataColumns: [
        { name: 'name', label: 'Name', align: 'left', field: 'Name', sortable: true },
        { name: 'email', label: 'Email', align: 'left', field: 'Email', sortable: true },
        { name: 'isAdmin', label: 'Admin', align: 'center', field: 'IsAdmin', sortable: true },

        { name: 'dailyTargetCall', label: 'Daily Target Call', align: 'center', field: row => row.IsAdmin? "" : row.Target.Daily.Call, sortable: true },
        { name: 'dailyTargetClosing', label: 'Daily Target Closing', align: 'center', field: row => row.IsAdmin? "" : row.Target.Daily.Closing, sortable: true },
        { name: 'dailyTargetBuyAmount', label: 'Daily Target Deposit Amount', align: 'center', field: row => row.IsAdmin? "" : row.Target.Daily.BuyAmount, sortable: true, format: (val, row) => row.IsAdmin? `` : `Rp. ${this.moneyFormat(val)}` },
        { name: 'weeklyTargetCall', label: 'Weekly Target Call', align: 'center', field: row => row.IsAdmin? "" : row.Target.Weekly.Call, sortable: true },
        { name: 'weeklyTargetClosing', label: 'Weekly Target Closing', align: 'center', field: row => row.IsAdmin? "" : row.Target.Weekly.Closing, sortable: true },
        { name: 'weeklyTargetBuyAmount', label: 'Weekly Target Deposit Amount', align: 'center', field: row => row.IsAdmin? "" : row.Target.Weekly.BuyAmount, sortable: true, format: (val, row) => row.IsAdmin? `` : `Rp. ${this.moneyFormat(val)}` },
        { name: 'monthlyTargetCall', label: 'Monthly Target Call', align: 'center', field: row => row.IsAdmin? "" : row.Target.Monthly.Call, sortable: true },
        { name: 'monthlyTargetClosing', label: 'Monthly Target Closing', align: 'center', field: row => row.IsAdmin? "" : row.Target.Monthly.Closing, sortable: true },
        { name: 'monthlyTargetBuyAmount', label: 'Monthly Target Deposit Amount', align: 'center', field: row => row.IsAdmin? "" : row.Target.Monthly.BuyAmount, sortable: true, format: (val, row) => row.IsAdmin? `` : `Rp. ${this.moneyFormat(val)}` },
        
        { name: 'dailyPerformanceCall', label: 'Daily Performance Call', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Daily.Call, sortable: true },
        { name: 'dailyPerformanceClosing', label: 'Daily Performance Closing', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Daily.Closing, sortable: true },
        { name: 'dailyPerformanceBuyAmount', label: 'Daily Performance Deposit Amount', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Daily.BuyAmount, sortable: true, format: (val, row) => row.IsAdmin? `` : `Rp. ${this.moneyFormat(val)}` },
        { name: 'weeklyPerformanceCall', label: 'Weekly Performance Call', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Weekly.Call, sortable: true },
        { name: 'weeklyPerformanceClosing', label: 'Weekly Performance Closing', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Weekly.Closing, sortable: true },
        { name: 'weeklyPerformanceBuyAmount', label: 'Weekly Performance Deposit Amount', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Weekly.BuyAmount, sortable: true, format: (val, row) => row.IsAdmin? `` : `Rp. ${this.moneyFormat(val)}` },
        { name: 'monthlyPerformanceCall', label: 'Monthly Performance Call', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Monthly.Call, sortable: true },
        { name: 'monthlyPerformanceClosing', label: 'Monthly Performance Closing', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Monthly.Closing, sortable: true },
        { name: 'monthlyPerformanceBuyAmount', label: 'Monthly Performance Deposit Amount', align: 'center', field: row => row.IsAdmin? "" : row.Performance.Monthly.BuyAmount, sortable: true, format: (val, row) => row.IsAdmin? `` : `Rp. ${this.moneyFormat(val)}` },

        { name: "action", align: "center", label: "Action" },
      ],
      telemarketerDataFilter: "",
      telemarketerDataVisible: ['name', 'email', 'isAdmin', 'monthlyPerformanceCall', 'monthlyPerformanceClosing', 'monthlyPerformanceBuyAmount', 'action'],
      telemarketerDataPagination: {
        rowsPerPage: 5 // current rows per page being displayed
      },
    }
  },

  mounted() {
    var vm = this;
    var data_submit = {
      Token: vm.$authService.getToken(),
      Limit: 10000,
    }
    this.$axios
      .post("/api/telemarketer/get", data_submit)
      .then(function (response) {
        if (response.data) {
          vm.telemarketers = response.data.Telemarketers
        }
      })
  },

  methods: {
    moneyFormat(value){
      const result = value.toString().replace(/\D/g, "").replace(/\B(?=(\d{3})+(?!\d))/g, ".")
      return result
    },
    onRowClick(evt, row) {
      console.log("clicked on", row.Email);
    },
    editTelemarketer(telemarketerID) {
      this.$router.push({
        name: "telemarketer-edit",
        params: { id: telemarketerID },
      });
    },
    confirmDeleteTelemarketer(telemarketer) {
      var vm = this;
      this.$q.dialog({
        title: 'Delete',
        message: 'Are you sure you want to delete ' + telemarketer.Email,
        cancel: true,
        persistent: true
      }).onOk(data => {
        vm.deleteTelemarketer(telemarketer.ID)
      }).onCancel(() => {
        // console.log('>>>> Cancel')
      }).onDismiss(() => {
        // console.log('I am triggered on both OK and Cancel')
      })      
    },
    deleteTelemarketer(telemarketerID) {
      var vm = this;
      var data_submit = {
        Telemarketer: {
          ID: telemarketerID
        }
      }
      this.$axios
      .post("/api/telemarketer/delete", data_submit)
      .then(function (response) {
        if (response.data) {
          vm.telemarketers.splice(vm.telemarketers.findIndex(item => item.ID === telemarketerID), 1)
        }
      })
    }
  }
}
</script>