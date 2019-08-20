<template lang="pug">
.column.is-full
  .card
    header.card-header
      p.card-header-title
        | {{title}} | {{courseId}}
    .card-content
      .content
        p level: {{level}}
        p Department : 
          nuxt-link(:to='"/dept/" + deptObjId ') {{deptName}} ({{deptId}})
        p Total Offers Count : {{Offers.length}}
        
        

      b-collapse(
        :open='false' 
        :aria-id='"contentId" + courseId ')
        button.button.is-primary(
          slot='trigger' 
          :aria-controls='"contentId" + courseId '
          ) more
        
        .notification
          .content
            p Created at : {{createdAt|timeFormat}}
            p Last updated : {{lastUpdated | timeFormat}}
            nuxt-link.button.is-info(:to='"/dept/" + deptObjId + "/course/" + objId ') Course Detail
            b-button.is-danger() Delete 
        
  
        //- .notification
        //-   .content
        //-     h4  Offers in over years
        //-     b-table(grouped group-multiline 
        //-       :data="Offers"
        //-       ref="table"
        //-       :opened-detailed="defaultOpenedDetails"
        //-       detailed
        //-       detail-key="_id"
        //-       :show-detail-icon="true"
        //-       :loading="false"
        //-       :paginated="isPaginated"
        //-       :per-page="perPage"
        //-       :current-page.sync="currentPage"
        //-       :pagination-simple="isPaginationSimple"
        //-       :pagination-position="paginationPosition"
        //-       :default-sort-direction="defaultSortDirection"
        //-       :sort-icon="sortIcon"
        //-       :sort-icon-size="sortIconSize"
        //-       default-sort="year"
        //-       aria-next-label="Next page"
        //-       aria-previous-label="Previous page"
        //-       aria-page-label="Page"
        //-       aria-current-label="Current page" )
              
        //-       template(slot-scope='props')
        //-         b-table-column(field='id' label='ID' width='40' :visible="false" )
        //-           | {{ props.row._id }}
        //-         b-table-column(field='year' label='Year' sortable='')
        //-           template(v-if='showDetailIcon')
        //-             | {{ props.row.year }}
        //-           template(v-else='')
        //-             a(@click='toggle(props.row)')
        //-               | {{ props.row.year }}

        //-         b-table-column(field='class_size' label='Class Size' sortable='')
        //-           | {{ props.row.class_size }}
        //-         b-table-column(field='available_places' label='Available Places' sortable='')
        //-           | {{ props.row.available_places }}

        //-       template(slot='detail' slot-scope='props')
        //-         .content
        //-           p Created at : {{props.row.created_at | timeFormat}}
        //-           p Last Update : {{props.row.updated_at | timeFormat}}
        //-           .columns
        //-             .column 
        //-               b-button(@click="editOffer(props.row)") Edit
        //-               b-button.is-danger(@click="delOffer(props.row)")  delete
                    
              

</template>

<script>
import moment from 'moment'
export default {
  data: () => ({
    // isPaginated: true,
    // isPaginationSimple: false,
    // paginationPosition: 'bottom',
    // defaultSortDirection: 'asc',
    // sortIcon: 'arrow-up',
    // sortIconSize: 'is-small',
    // currentPage: 1,
    // perPage: 10,
    // defaultOpenedDetails: [1],
  }),
  props: {
    objId :{ type: String },
    deptId : {type : String},
    deptName : { type  : String },
    deptObjId : { type : String },
    courseId:{ type : String },
    title : { type : String },
    level : { type : Number },
    offerCount :{ type : Number },
    Offers : { type : Array },
    createdAt : { type : String },
    lastUpdated : { type :String },
  },
  computed : {
    toLink : function () {
      return '/course/' + this.objId
    },
    latestOffer : function () {
        return this.Offers.sort((a,b)=>{
          return b.year - a.year
        })
    }
  },
  filters: {
    timeFormat (i){
      return moment(i , "YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss")
    }
  },
  methods: {
    toggle(row) {
      this.$refs.table.toggleDetails(row)
    },
    editOffer(row) {

    }
  },
  mounted() {
    console.log(this.lastUpdated)
    console.log(this.createdAt)
    console.log(this.latestOffer)
  }
}
</script>
