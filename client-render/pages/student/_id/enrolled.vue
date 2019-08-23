<template lang="pug">
div
  section.hero.is-info
    .hero-body
      .container.has-text-left
        p.title {{student_name}}
        p.subtitle {{student_id}}
    .hero-foot
      nav.tabs.is-boxed.is-fullwidth
        .container
          ul
            li
              nuxt-link(:to='"/student/" + this.objId ')  Overview
            li.is-active
              nuxt-link(:to='"/student/" + this.objId + "/enrolled" ') Enrolled Course List

  section.section
    h2.title  Enrolled Course
    .container
      .columns.is-multiline
        .column
          b-collapse(
            :open.sync='createBlock_isopen'
            aria-id='createEnroll')
            b-button.is-primary(
              @click="createReady()"
              slot='trigger' 
              aria-controls='createEnroll') Enroll New Course
            section.section
              b-field(grouped group-multiline)
                b-field(label='Course' horizontal )
                  b-select(placeholder='Select a Course' width="160" v-model='co_course_id')
                    option(v-for='option in createOptCourse' :value='option.course_id' :key='option._id')
                      | {{ option.course_id }} | {{option.title}}
              
                  //- b-field(grouped group-multiline)
                b-field(label='Offer' horizontal)
                  b-select(placeholder='Offer Year' v-model='co_year')
                    option(v-for='option in createOptYear' :value='option.offers.year' :key='option.offers._id')
                      | {{ option.offers.year }} 

                b-field(label='Enroll Date' horizontal)
                  b-datepicker( v-model='co_date' )
                  
                  //- b-field(grouped group-multiline)
                b-field
                  b-button.is-primary(@click="createEnroll()") Create
                b-field
                  b-button(@click="cancelCreate()") Cancel
                
      .content  
        b-table(grouped group-multiline 
          :data="enrolled"
          ref="table"
          :opened-detailed="defaultOpenedDetails"
          detailed
          detail-key="_id"
          :show-detail-icon="false"
          :loading="false"
          :paginated="isPaginated"
          :per-page="perPage"
          :current-page.sync="currentPage"
          :pagination-simple="isPaginationSimple"
          :pagination-position="paginationPosition"
          :default-sort-direction="defaultSortDirection"
          :sort-icon="sortIcon"
          :sort-icon-size="sortIconSize"
          default-sort="year"
          aria-next-label="Next page"
          aria-previous-label="Previous page"
          aria-page-label="Page"
          aria-current-label="Current page" )
          
          template(slot-scope='props')
            b-table-column(field='id' label='ID' width='40' :visible="false" )
              | {{ props.row._id }}
            b-table-column(field='course' label='Course' width='160' sortable='')
              | {{ props.row.course_id }}
            b-table-column(field='year' label='Year' width='100' sortable='')
              | {{ props.row.year }}
            b-table-column(field='enroll_date' label='Enroll Date'  width='160' sortable='')
              | {{ props.row.enroll_date | timeFormat }}
            b-table-column(field='created_at' label='Created at'  width='160' sortable='')
              | {{props.row.created_at | timeFormat}}
            b-table-column(field='updated_at' label='Last Updated at'  width='160' sortable='')
              | {{props.row.updated_at | timeFormat}}
            b-table-column(field='edit' label='Edit' width='70')
              b-button(@click="toggle(props.row)") Edit
            b-table-column(field='delete' label='Delete' width='70')
              b-button.is-danger(@click="warn(props.row)") Delete


          template(slot='detail' slot-scope='props')
            section
              b-field(grouped group-multiline)
                b-field(label='Course' horizontal )
                  b-select(placeholder='Select a Course' width="160" v-model='tmpRow.course_id')
                    option(v-for='option in createOptCourse' :value='option.course_id' :key='option._id')
                      | {{ option.course_id }} | {{option.title}}
              
                b-field(label='Offer' horizontal)
                  b-select(placeholder='Offer Year' v-model='tmpRow.year')
                    option(v-for='option in editOptYear' :value='option.offers.year' :key='option.offers._id')
                      | {{ option.offers.year }} 

                b-field(label='Enroll Date' horizontal)
                  b-datepicker( v-model='tmpRow.enroll_date' )
                b-field
                  b-button.is-primary(@click="saveEnroll()") Save
                b-field
                  b-button(@click="closeCancel(props.row)") Cancel
</template>

<script>
import moment from "moment";
import cdw from "lodash/cloneDeep";
export default {
  name: "OfferPage",
  data: () => ({
    student_name: "",
    student_id: "",
    objId: "",
    enrolled: [],

    OfferOpt: [],

    isPaginated: true,
    isPaginationSimple: false,
    paginationPosition: "bottom",
    defaultSortDirection: "desc",
    sortIcon: "arrow-up",
    sortIconSize: "is-small",
    currentPage: 1,
    perPage: 10,
    defaultOpenedDetails: [],

    tmpRow: {},

    co_course_id: "",
    co_year: new Date().getFullYear(),
    co_date: new Date(),

    createBlock_isopen: false
  }),
  filters: {
    timeFormat(i) {
      return moment(i, "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss");
    }
  },
  computed: {
    createOptCourse() {
      return this.OfferOpt
        ? this.OfferOpt.reduce((i, c) => {
            if (i.length === 0 || i[i.length - 1].course_id !== c.course_id) {
              i.push(c);
            }
            return i;
          }, [])
        : [];
    },
    createOptYear() {
      return this.OfferOpt
        ? this.OfferOpt.filter(e => {
            return e.course_id == this.co_course_id;
          })
        : [];
    },
    editOptYear() {
      return this.OfferOpt
        ? this.OfferOpt.filter(e => {
            return e.course_id == this.tmpRow.course_id;
          })
        : [];
    }
  },
  methods: {
    async fetchStudent() {
      try {
        let ip = await this.$axios.$get(
          "/api/v1/g/student/" + this.$route.params["id"]
        );
        this.student_name = ip.data.student_name;
        this.student_id = ip.data.student_id;
        this.objId = ip.data._id;
        this.enrolled = ip.data.enrolled;
      } catch (e) {
        this.$buefy.toast.open({
          message: "The Requested Student is not exist",
          type: "is-warning",
          position: "is-bottom"
        });
        this.$router.push({
          path: "/student/" + this.$route.params["id"]
        });
      }
    },

    async fetchOffer() {
      try {
        let ip = await this.$axios.$get(
          `/api/v1/l/course/_/offer?available_places={"gt":0}`
        );
        this.OfferOpt = ip.data;
        console.log(ip);
      } catch (e) {
        this.$buefy.toast.open({
          message: "The Offer List fetching error",
          type: "is-warning",
          position: "is-bottom"
        });
      }
    },

    toggle(row) {
      this.createBlock_isopen = false;
      this.tmpRow = cdw(row);
      this.tmpRow["enroll_date"] = moment(this.tmpRow["enroll_date"]).toDate()
      this.defaultOpenedDetails.pop();
      this.$refs.table.toggleDetails(row);
    },

    async saveEnroll() {
      console.log(this.tmpRow);
      try {
        let ip = await this.$axios.$post(
          "/api/v1/u/student/" +
            this.objId +
            "/enrolled/" +
            this.tmpRow._id,
          {
            year: Number(this.tmpRow.year),
            enroll_date: String(moment(this.tmpRow.enroll_date).format("YYYY-MM-DDTHH:mm:ssZ")),
            course_id: String(this.tmpRow.course_id)
          }
        );
        console.log(ip);
        this.$buefy.toast.open({
          duration: 5000,
          message: `Save Success !`,
          position: "is-bottom",
          type: "is-success"
        });
      } catch (e) {
        this.$buefy.toast.open({
          duration: 5000,
          message: `Save Fail !`,
          position: "is-bottom",
          type: "is-danger"
        });
      } finally {
        this.fetchStudent();
      }
    },

    warn(row) {
      console.log(row);
      this.$buefy.dialog.confirm({
        title: "Deleting Offer",
        message:
          "Are you sure you want to <b>delete</b> this Enroll Record? This action cannot be undone.",
        confirmText: "Delete Enroll Record",
        type: "is-danger",
        hasIcon: true,
        onConfirm:()=> this.deleteEnroll(row),
      });
    },
    async deleteEnroll(row) {
      try {
        let ip = await this.$axios.$post(
          "/api/v1/d/student/" + this.objId + "/enrolled/" + row._id,
          {
            _id: String(row._id)
          }
        );
        this.$buefy.toast.open({
          duration: 5000,
          message: `deleted !`,
          position: "is-bottom",
          type: "is-success"
        });
      } catch (e) {
        console.log(e);
        this.$buefy.toast.open({
          duration: 5000,
          message: `Something's not good`,
          position: "is-bottom",
          type: "is-danger"
        });
      } finally {
        this.fetchStudent();
      }
    },
    // Create block
    createReady() {
      let lastOffer = this.offers
        ? this.offer.sort((a, b) => {
            return b.year - a.year;
          })
        : [];
      this.defaultOpenedDetails.pop();
    },
    async createEnroll() {
      let ytmpRow = {
        year: Number(this.co_year),
        course_id: String(this.co_course_id),
        enroll_date: String(moment(this.co_date).format("YYYY-MM-DDTHH:mm:ssZ"))
      };
      console.log(ytmpRow)
      try {
        let ip = await this.$axios.$post(
          "/api/v1/c/student/" + this.objId + "/enrolled",
          ytmpRow
        );
        this.$buefy.toast.open({
          message: "Offer Create Success",
          type: "is-success",
          position: "is-bottom"
        });
      } catch (e) {
        console.warn(e);
        this.$buefy.toast.open({
          message: "Offer Create Fail",
          type: "is-danger",
          position: "is-bottom"
        });
      } finally {
        this.fetchStudent();
      }
    },
    cancelCreate() {
      this.createBlock_isopen = false;
    }
  },
  beforeMount() {
    this.fetchStudent();
    this.fetchOffer();
  }
};
</script>
