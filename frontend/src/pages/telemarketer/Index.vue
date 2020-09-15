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
          <q-input color="grey-8" dense debounce="300" v-model="telemarketerDataFilter" placeholder="Search">
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
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
        { name: 'weeklyTargetCall', label: 'Weekly Target Call', align: 'center', field: 'WeeklyTarget.Call', sortable: true },
        { name: 'weeklyTargetClosing', label: 'Weekly Target Closing', align: 'center', field: 'WeeklyTarget.Closing', sortable: true },
        { name: 'performance', label: 'Performance', align: 'center', field: 'Performance', sortable: true },
        { name: "action", align: "center", label: "Action" },
      ],
      telemarketerDataFilter: "",
      telemarketerDataVisible: ['name', 'email', 'isAdmin', 'weeklyTargetCall', 'weeklyTargetClosing', 'performance', 'action'],
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
    onRowClick(evt, row) {
      console.log("clicked on", row.Email);
    },
    editTelemarketer(telemarketerID) {
      this.$router.push({
        name: "telemarketer-edit",
        params: { id: telemarketerID },
      });
    },
  }
}
</script>