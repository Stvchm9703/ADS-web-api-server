<template lang="pug">
div
  section.hero.is-info
    .hero-body
      .container.has-text-left
        p.title.is-left {{course_title}} ({{course_id}})
        p.subtitle {{deptName}} ({{deptId}})
    .hero-foot
      nav.tabs.is-boxed.is-fullwidth
        .container
          ul
            li
              nuxt-link(:to='"/dept/" + deptObjId+ "/course/" + course_obj_id ')  Overview
            li.is-active
              nuxt-link(:to='"/dept/" + deptObjId+ "/course/" + course_obj_id  + "/offer" ')  Offers List
            li


  section.section
    h2.title  Offers in over years
    .container
      .columns.is-multiline
        .column
          b-collapse(
            :open.sync='createBlock_isopen'
            aria-id='createOfferId')
            b-button.is-primary(
              @click="createReady()"
              slot='trigger' 
              aria-controls='createOfferId') Create Offer
            section.section
              b-field(grouped group-multiline)
                b-field( label="Year"  horizontal)
                  b-input(
                    type="number"
                    placeholder="Year" 
                    v-model="tmpRow.year")
                b-field( label="Class size"  horizontal )
                  b-input(
                    type="number"
                    placeholder="Class size"
                    v-model="tmpRow.class_size")
                b-field( label="Available Places"  horizontal)
                  b-input(
                    type="number"
                    placeholder="Available Places"
                    width="40"
                    v-model="tmpRow.available_places")
              b-field(grouped group-multiline)
                b-field
                  b-button.is-primary(@click="createOffer()") Create
                b-field
                  b-button(@click="cancelCreate()") Cancel
                
      .content  
        b-table(grouped group-multiline 
          :data="offers"
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
            b-table-column(field='year' label='Year' width='100' sortable='')
              | {{ props.row.year }}
            b-table-column(field='class_size' label='Class Size' width='160' sortable='')
              | {{ props.row.class_size }}
            b-table-column(field='available_places' label='Available Places'  width='160' sortable='')
              | {{ props.row.available_places }}
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
                b-field( label="Year"  horizontal)
                  b-input(placeholder="Year" 
                    type="number"  v-model="tmpRow.year")

                b-field( label="Class size"  horizontal )
                  b-input(placeholder="Class size"
                    type="number"  v-model="tmpRow.class_size")
                    
                b-field( label="Available Places"  horizontal)
                  b-input(placeholder="Available Places"
                    width="40"
                    type="number"  v-model="tmpRow.available_places")

                b-field
                  b-button(@click="saveOffer()") Save
                b-field
                  b-button(@click="closeCancel(props.row)") Cancel
</template>

<script>
import moment from "moment";
import cdw from "lodash/cloneDeep";
export default {
  name: "OfferPage",
  data: () => ({
    deptName: "",
    deptId: "",
    deptObjId: "",
    course_obj_id: "",
    course_id: "",
    course_title: "",
    lastUpdated: "",
    createdAt: "",
    offers: [],

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

    createBlock_isopen: false
  }),
  filters: {
    timeFormat(i) {
      return moment(i, "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss");
    }
  },
  methods: {
    async fetchDept() {
      let ip = await this.$axios.$get(
        "/api/v1/g/dept/" + this.$route.params["id"]
      );
      this.deptName = ip.data.dept_name;
      this.deptId = ip.data.dept_id;
      this.deptObjId = ip.data._id;
    },

    async fetchCourse() {
      try {
        let ip = await this.$axios.$get(
          "/api/v1/g/dept/" +
            this.$route.params["id"] +
            "/course/" +
            this.$route.params["course_id"]
        );
        console.log(ip);
        this.course_obj_id = ip.data._id;
        this.course_id = ip.data.course_id;
        this.course_title = ip.data.title;
        this.lastUpdated = ip.data.updated_at;
        this.createdAt = ip.data.created_at;
        this.offers = ip.data.offers;
      } catch (e) {
        // console.log(e);
        // console.log("the requested dept is not exists");
        // console.log("redirect to home");
        this.$buefy.toast.open({
          message: "The Requested Department is not exist",
          type: "is-warning",
          position: "is-bottom"
        });
        this.$router.push({
          path: "/dept/" + this.$route.params["id"]
        });
      }
    },

    async fetchOffer() {
      let ip = await this.$axios.$get(
        "/api/v1/l/course/" + this.$route.params["course_id"] + "/offer"
      );
      console.log(ip);
    },

    toggle(row) {
      this.createBlock_isopen = false;
      this.tmpRow = cdw(row);
      this.defaultOpenedDetails.pop();
      this.$refs.table.toggleDetails(row);
    },

    async saveOffer() {
      console.log(this.tmpRow);
      try {
        let ip = await this.$axios.$post(
          "/api/v1/u/course/" +
            this.course_obj_id +
            "/offer/" +
            this.tmpRow._id,
          {
            _id: String(this.tmpRow._id),
            year: Number(this.tmpRow.year),
            available_places: Number(this.tmpRow.available_places),
            class_size: Number(this.tmpRow.class_size)
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
        this.fetchCourse();
      }
    },

    warn(row) {
      this.$buefy.dialog.confirm({
        title: "Deleting Offer",
        message:
          "Are you sure you want to <b>delete</b> this Offer? This action cannot be undone.",
        confirmText: "Delete Offer",
        type: "is-danger",
        hasIcon: true,
        onConfirm: () =>this.deleteOffer(row)
      });
    },
    async deleteOffer(row) {
      try {
        let ip = await this.$axios.$post(
          "/api/v1/d/course/" + this.course_obj_id + "/offer/" + row._id,
          { _id: String(row._id) }
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
        this.fetchCourse();
      }
    },
    // Create block
    createReady() {
      let lastOffer = this.offers 
        ? this.offers.sort((a, b) => {
            return b.year - a.year;
          })
        : [];
      let dt = new Date();
      this.tmpRow = {
        year: dt.getFullYear(),
        class_size: this.offers ? lastOffer[0].class_size : 40,
        available_places: this.offers ? lastOffer[0].class_size : 40
      };
      console.log(dt.getFullYear());
      this.defaultOpenedDetails.pop();
    },
    async createOffer() {
      console.log(this.tmpRow);
      let ytmpRow = {
        year: Number(this.tmpRow.year),
        class_size: Number(this.tmpRow.class_size),
        available_places: Number(this.tmpRow.available_places)
      };
      try {
        let ip = await this.$axios.$post(
          "/api/v1/c/course/" + this.course_obj_id + "/offer",
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
        this.fetchCourse();
      }
    },
    cancelCreate() {
      this.createBlock_isopen = false;
    }
  },
  beforeMount() {
    this.fetchDept();
    this.fetchCourse();
  }
};
</script>
