(window.webpackJsonp=window.webpackJsonp||[]).push([[17],{410:function(t,e,n){var map={"./af":279,"./af.js":279,"./ar":280,"./ar-dz":281,"./ar-dz.js":281,"./ar-kw":282,"./ar-kw.js":282,"./ar-ly":283,"./ar-ly.js":283,"./ar-ma":284,"./ar-ma.js":284,"./ar-sa":285,"./ar-sa.js":285,"./ar-tn":286,"./ar-tn.js":286,"./ar.js":280,"./az":287,"./az.js":287,"./be":288,"./be.js":288,"./bg":289,"./bg.js":289,"./bm":290,"./bm.js":290,"./bn":291,"./bn.js":291,"./bo":292,"./bo.js":292,"./br":293,"./br.js":293,"./bs":294,"./bs.js":294,"./ca":295,"./ca.js":295,"./cs":296,"./cs.js":296,"./cv":297,"./cv.js":297,"./cy":298,"./cy.js":298,"./da":299,"./da.js":299,"./de":300,"./de-at":301,"./de-at.js":301,"./de-ch":302,"./de-ch.js":302,"./de.js":300,"./dv":303,"./dv.js":303,"./el":304,"./el.js":304,"./en-SG":305,"./en-SG.js":305,"./en-au":306,"./en-au.js":306,"./en-ca":307,"./en-ca.js":307,"./en-gb":308,"./en-gb.js":308,"./en-ie":309,"./en-ie.js":309,"./en-il":310,"./en-il.js":310,"./en-nz":311,"./en-nz.js":311,"./eo":312,"./eo.js":312,"./es":313,"./es-do":314,"./es-do.js":314,"./es-us":315,"./es-us.js":315,"./es.js":313,"./et":316,"./et.js":316,"./eu":317,"./eu.js":317,"./fa":318,"./fa.js":318,"./fi":319,"./fi.js":319,"./fo":320,"./fo.js":320,"./fr":321,"./fr-ca":322,"./fr-ca.js":322,"./fr-ch":323,"./fr-ch.js":323,"./fr.js":321,"./fy":324,"./fy.js":324,"./ga":325,"./ga.js":325,"./gd":326,"./gd.js":326,"./gl":327,"./gl.js":327,"./gom-latn":328,"./gom-latn.js":328,"./gu":329,"./gu.js":329,"./he":330,"./he.js":330,"./hi":331,"./hi.js":331,"./hr":332,"./hr.js":332,"./hu":333,"./hu.js":333,"./hy-am":334,"./hy-am.js":334,"./id":335,"./id.js":335,"./is":336,"./is.js":336,"./it":337,"./it-ch":338,"./it-ch.js":338,"./it.js":337,"./ja":339,"./ja.js":339,"./jv":340,"./jv.js":340,"./ka":341,"./ka.js":341,"./kk":342,"./kk.js":342,"./km":343,"./km.js":343,"./kn":344,"./kn.js":344,"./ko":345,"./ko.js":345,"./ku":346,"./ku.js":346,"./ky":347,"./ky.js":347,"./lb":348,"./lb.js":348,"./lo":349,"./lo.js":349,"./lt":350,"./lt.js":350,"./lv":351,"./lv.js":351,"./me":352,"./me.js":352,"./mi":353,"./mi.js":353,"./mk":354,"./mk.js":354,"./ml":355,"./ml.js":355,"./mn":356,"./mn.js":356,"./mr":357,"./mr.js":357,"./ms":358,"./ms-my":359,"./ms-my.js":359,"./ms.js":358,"./mt":360,"./mt.js":360,"./my":361,"./my.js":361,"./nb":362,"./nb.js":362,"./ne":363,"./ne.js":363,"./nl":364,"./nl-be":365,"./nl-be.js":365,"./nl.js":364,"./nn":366,"./nn.js":366,"./pa-in":367,"./pa-in.js":367,"./pl":368,"./pl.js":368,"./pt":369,"./pt-br":370,"./pt-br.js":370,"./pt.js":369,"./ro":371,"./ro.js":371,"./ru":372,"./ru.js":372,"./sd":373,"./sd.js":373,"./se":374,"./se.js":374,"./si":375,"./si.js":375,"./sk":376,"./sk.js":376,"./sl":377,"./sl.js":377,"./sq":378,"./sq.js":378,"./sr":379,"./sr-cyrl":380,"./sr-cyrl.js":380,"./sr.js":379,"./ss":381,"./ss.js":381,"./sv":382,"./sv.js":382,"./sw":383,"./sw.js":383,"./ta":384,"./ta.js":384,"./te":385,"./te.js":385,"./tet":386,"./tet.js":386,"./tg":387,"./tg.js":387,"./th":388,"./th.js":388,"./tl-ph":389,"./tl-ph.js":389,"./tlh":390,"./tlh.js":390,"./tr":391,"./tr.js":391,"./tzl":392,"./tzl.js":392,"./tzm":393,"./tzm-latn":394,"./tzm-latn.js":394,"./tzm.js":393,"./ug-cn":395,"./ug-cn.js":395,"./uk":396,"./uk.js":396,"./ur":397,"./ur.js":397,"./uz":398,"./uz-latn":399,"./uz-latn.js":399,"./uz.js":398,"./vi":400,"./vi.js":400,"./x-pseudo":401,"./x-pseudo.js":401,"./yo":402,"./yo.js":402,"./zh-cn":403,"./zh-cn.js":403,"./zh-hk":404,"./zh-hk.js":404,"./zh-tw":405,"./zh-tw.js":405};function l(t){var e=r(t);return n(e)}function r(t){if(!n.o(map,t)){var e=new Error("Cannot find module '"+t+"'");throw e.code="MODULE_NOT_FOUND",e}return map[t]}l.keys=function(){return Object.keys(map)},l.resolve=r,t.exports=l,l.id=410},525:function(t,e,n){"use strict";n.r(e);n(33);var l,r,o=n(6),d=n(278),c=n.n(d),j={props:{objId:String,student_name:{type:String},student_id:{type:String},dob:String,enrolled:{type:Array,default:function(){return[]}},createdAt:String,lastUpdated:String},filters:{timeFormat:function(i){return c()(i,"YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD HH:mm:ss")},dobForm:function(i){return c()(i,"YYYY-MM-DDTHH:mm:ssZ").format("YYYY-MM-DD")}},methods:{warnDelete:function(){var t=this;this.$buefy.dialog.confirm({title:"Deleting Student",message:"Are you sure you want to <b>delete</b> this Student? This action cannot be undone.",confirmText:"Delete Student",type:"is-danger",hasIcon:!0,onConfirm:function(){t.deleteStudent()}})},deleteStudent:(l=Object(o.a)(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,this.$axios.$post("/api/v1/d/student/",{_id:String(this.objId)});case 2:e=t.sent,console.log(e),this.$buefy.toast.open({duration:5e3,message:"deleted !",position:"is-bottom",type:"is-success"}),this.$router.push({path:"/student"});case 6:case"end":return t.stop()}},t,this)})),function(){return l.apply(this,arguments)})}},f=n(15),h={components:{Info:Object(f.a)(j,function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",{staticClass:"container"},[n("h1",{staticClass:"title"},[t._v("Information")]),n("h2",{staticClass:"subtitle"},[t._v("Basic information of the Student")]),n("div",{staticClass:"columns"},[t._m(0),n("div",{staticClass:"column is-third-quarter-fullhd is-two-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[n("p",[t._v(t._s(t.student_name))])])]),n("div",{staticClass:"columns"},[t._m(1),n("div",{staticClass:"column is-third-quarter-fullhd is-two-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[n("p",[t._v(t._s(t.student_id))])])]),n("div",{staticClass:"columns"},[t._m(2),n("div",{staticClass:"column is-third-quarter-fullhd is-two-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[n("p",[t._v(t._s(t._f("dobForm")(t.dob)))])])]),n("div",{staticClass:"columns"},[t._m(3),n("div",{staticClass:"column is-third-quarter-fullhd is-two-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[n("p",[t._v(t._s(t.enrolled?t.enrolled.length:0))])])]),t._m(4),n("div",{staticClass:"columns"},[n("div",{staticClass:"column is-one-third-fullhd is-one-third-widescreen is-half-desktop is-full-tablet is-full-mobile"},[n("h3",[t._v("Created at : "+t._s(t._f("timeFormat")(t.createdAt))+"  ")])]),n("div",{staticClass:"column is-one-third-fullhd is-one-third-widescreen is-half-desktop is-full-tablet is-full-mobile"},[n("h3",[t._v("Last Update at :  "+t._s(t._f("timeFormat")(t.lastUpdated))+"  ")])])]),n("div",{staticClass:"columns is-vcentered"},[n("div",{staticClass:"column is-half-fullhd is-half-widescreen is-half-desktop is-full-tablet is-full-mobile"},[n("nuxt-link",{staticClass:"button is-primary",attrs:{to:"/student/"+t.objId+"/update"}},[t._v("  Update information")])],1),n("div",{staticClass:"column is-half-fullhd is-half-widescreen is-half-desktop is-full-tablet is-full-mobile"},[n("b-button",{staticClass:"is-danger",on:{click:function(e){return t.warnDelete()}}},[t._v(" Delete information")])],1)])])},[function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"column is-one-quarter-fullhd is-one-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[e("h3",[this._v(" Student name ")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"column is-one-quarter-fullhd is-one-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[e("h3",[this._v(" Student Id Code ")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"column is-one-quarter-fullhd is-one-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[e("h3",[this._v(" Day of Birth ")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"column is-one-quarter-fullhd is-one-third-widescreen is-full-desktop is-full-tablet is-full-mobile"},[e("h3",[this._v(" Total Enrolled Count: ")])])},function(){var t=this.$createElement,e=this._self._c||t;return e("div",{staticClass:"columns"},[e("div",{staticClass:"column is-full"},[e("hr")])])}],!1,null,null,null).exports},data:function(){return{objId:"",student_name:"",student_id:"",enrolled:[],dob:"",lastUpdated:"",createdAt:""}},methods:{fetchDept:(r=Object(o.a)(regeneratorRuntime.mark(function t(){var e;return regeneratorRuntime.wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,this.$axios.$get("/api/v1/g/student/"+this.$route.params.id);case 2:e=t.sent,this.objId=e.data._id,this.student_name=e.data.student_name,this.student_id=e.data.student_id,this.enrolled=null==e.data.enrolled?[]:e.data.enrolled,this.dob=e.data.dob,this.lastUpdated=e.data.updated_at,this.createdAt=e.data.created_at;case 10:case"end":return t.stop()}},t,this)})),function(){return r.apply(this,arguments)})},beforeMount:function(){this.fetchDept()}},m=Object(f.a)(h,function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("div",[n("section",{staticClass:"hero is-info"},[n("div",{staticClass:"hero-body"},[n("div",{staticClass:"container has-text-left"},[n("p",{staticClass:"title"},[t._v(t._s(t.student_name))]),n("p",{staticClass:"subtitle"},[t._v(t._s(t.student_id))])])]),n("div",{staticClass:"hero-foot"},[n("nav",{staticClass:"tabs is-boxed is-fullwidth"},[n("div",{staticClass:"container"},[n("ul",[n("li",{staticClass:"is-active"},[n("nuxt-link",{attrs:{to:"/student/"+this.objId}},[t._v(" Overview")])],1),n("li",[n("nuxt-link",{attrs:{to:"/student/"+this.objId+"/enrolled"}},[t._v("Enrolled Course List")])],1)])])])])]),n("section",{staticClass:"section"},[n("Info",{attrs:{objId:t.objId,student_name:t.student_name,student_id:t.student_id,dob:t.dob,enrolled:t.enrolled,createdAt:t.createdAt,lastUpdated:t.lastUpdated}})],1)])},[],!1,null,null,null);e.default=m.exports}}]);