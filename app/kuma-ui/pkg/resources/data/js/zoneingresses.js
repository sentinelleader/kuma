"use strict";(self["webpackChunkkuma_gui"]=self["webpackChunkkuma_gui"]||[]).push([[639],{84855:function(t,e,n){n.d(e,{Z:function(){return l}});var s=function(){var t=this,e=t._self._c;return e("div",{staticClass:"component-frame"},[t._t("default")],2)},i=[],a={name:"FrameSkeleton"},r=a,o=n(1001),u=(0,o.Z)(r,s,i,!1,null,"666bca0e",null),l=u.exports},17044:function(t,e,n){n.r(e),n.d(e,{default:function(){return z}});var s=function(){var t=this,e=t._self._c;return e("div",{staticClass:"zoneingresses"},[!1===t.multicluster?e("MultizoneInfo"):e("FrameSkeleton",[e("DataOverview",{attrs:{"page-size":t.pageSize,"has-error":t.hasError,"is-loading":t.isLoading,"empty-state":t.empty_state,"table-data":t.tableData,"table-data-is-empty":t.isEmpty,next:t.next},on:{tableAction:t.tableAction,loadData:function(e){return t.loadData(e)}},scopedSlots:t._u([{key:"additionalControls",fn:function(){return[t.$route.query.ns?e("KButton",{staticClass:"back-button",attrs:{appearance:"primary",size:"small",to:{name:"zoneingresses"}}},[e("span",{staticClass:"custom-control-icon"},[t._v(" ← ")]),t._v(" View All ")]):t._e()]},proxy:!0}])}),!1===t.isEmpty?e("TabsWidget",{attrs:{"has-error":t.hasError,"is-loading":t.isLoading,tabs:t.tabs,"initial-tab-override":"overview"},scopedSlots:t._u([{key:"tabHeader",fn:function(){return[e("div",[e("h3",[t._v(" Zone Ingress: "+t._s(t.entity.name))])]),e("div",[e("EntityURLControl",{attrs:{name:t.entity.name}})],1)]},proxy:!0},{key:"overview",fn:function(){return[e("LabelList",[e("div",[e("ul",t._l(t.entity,(function(n,s){return e("li",{key:s},[e("h4",[t._v(" "+t._s(s)+" ")]),e("p",[t._v(" "+t._s(n)+" ")])])})),0)])])]},proxy:!0},{key:"insights",fn:function(){return[e("KCard",{attrs:{"border-variant":"noBorder"},scopedSlots:t._u([{key:"body",fn:function(){return[e("AccordionList",{attrs:{"initially-open":0}},t._l(t.subscriptionsReversed,(function(n,s){return e("AccordionItem",{key:s,scopedSlots:t._u([{key:"accordion-header",fn:function(){return[e("SubscriptionHeader",{attrs:{details:n}})]},proxy:!0},{key:"accordion-content",fn:function(){return[e("SubscriptionDetails",{attrs:{details:n,"is-discovery-subscription":""}})]},proxy:!0}],null,!0)})})),1)]},proxy:!0}],null,!1,4118320068)})]},proxy:!0},{key:"xds-configuration",fn:function(){return[e("XdsConfiguration",{attrs:{"zone-ingress-name":t.entity.name}})]},proxy:!0},{key:"envoy-stats",fn:function(){return[e("EnvoyStats",{attrs:{"zone-ingress-name":t.entity.name}})]},proxy:!0},{key:"envoy-clusters",fn:function(){return[e("EnvoyClusters",{attrs:{"zone-ingress-name":t.entity.name}})]},proxy:!0}],null,!1,3591657575)}):t._e()],1)],1)},i=[],a=n(27361),r=n.n(a),o=n(20629),u=n(99716),l=n(4104),c=n(58641),d=n(70172),y=n(53419),p=n(17463),m=n(84855),f=n(56882),g=n(87673),h=n(7001),v=n(33561),b=n(66190),_=n(65404),k=n(45689),Z=n(74473),S=n(49718),x=n(64082),E=n(46077),D={name:"ZoneIngresses",components:{EnvoyClusters:E.Z,EnvoyStats:x.Z,FrameSkeleton:m.Z,DataOverview:f.Z,TabsWidget:h.Z,LabelList:v.Z,AccordionList:Z.Z,AccordionItem:S.Z,SubscriptionDetails:u.Z,SubscriptionHeader:l.Z,MultizoneInfo:c.Z,EntityURLControl:g.Z,XdsConfiguration:b.Z},metaInfo:{title:"ZoneIngresses"},data(){return{isLoading:!0,isEmpty:!1,hasError:!1,empty_state:{title:"No Data",message:"There are no Zone Ingresses present."},tableData:{headers:[{key:"actions",hideLabel:!0},{label:"Status",key:"status"},{label:"Name",key:"name"}],data:[]},tabs:[{hash:"#overview",title:"Overview"},{hash:"#insights",title:"Zone Ingress Insights"},{hash:"#xds-configuration",title:"XDS Configuration"},{hash:"#envoy-stats",title:"Stats"},{hash:"#envoy-clusters",title:"Clusters"}],entity:{},rawData:[],pageSize:k.NR,next:null,subscriptionsReversed:[]}},computed:{...(0,o.Se)({multicluster:"config/getMulticlusterStatus"})},watch:{$route(){this.init()}},beforeMount(){this.init()},methods:{init(){this.multicluster&&this.loadData()},tableAction(t){const e=t;this.getEntity(e)},async loadData(t="0"){this.isLoading=!0,this.isEmpty=!1;const e=this.$route.query.ns||null;try{const{data:n,next:s}=await(0,d.W)({getAllEntities:p.Z.getAllZoneIngressOverviews.bind(p.Z),getSingleEntity:p.Z.getZoneIngressOverview.bind(p.Z),size:this.pageSize,offset:t,query:e});this.next=s,n.length?(this.isEmpty=!1,this.rawData=n,this.getEntity({name:n[0].name}),this.tableData.data=n.map((t=>{const{zoneIngressInsight:e={}}=t;return{...t,...(0,_._I)(e)}}))):(this.tableData.data=[],this.isEmpty=!0)}catch(n){this.hasError=!0,this.isEmpty=!0,console.error(n)}finally{this.isLoading=!1}},getEntity(t){const e=["type","name"],n=this.rawData.find((e=>e.name===t.name)),s=r()(n,"zoneIngressInsight.subscriptions",[]);this.subscriptionsReversed=Array.from(s).reverse(),this.entity=(0,y.wy)(n,e)}}},I=D,C=n(1001),w=(0,C.Z)(I,s,i,!1,null,null,null),z=w.exports},58641:function(t,e,n){n.d(e,{Z:function(){return d}});var s=function(){var t=this,e=t._self._c;return e("KEmptyState",{scopedSlots:t._u([{key:"title",fn:function(){return[e("KIcon",{staticClass:"kong-icon--centered",attrs:{icon:"dangerCircle",size:"42"}}),t._v(" "+t._s(t.productName)+" is running in Standalone mode. ")]},proxy:!0},{key:"message",fn:function(){return[e("p",[t._v(" To access this page, you must be running in "),e("strong",[t._v("Multi-Zone")]),t._v(" mode. ")])]},proxy:!0},{key:"cta",fn:function(){return[e("KButton",{attrs:{to:`https://kuma.io/docs/${t.kumaDocsVersion}/documentation/deployments/`,target:"_blank",appearance:"primary"}},[t._v(" Learn More ")])]},proxy:!0}])})},i=[],a=n(20629),r=n(45689),o={name:"MultizoneInfo",data(){return{productName:r.sG}},computed:{...(0,a.Se)({kumaDocsVersion:"config/getKumaDocsVersion"})}},u=o,l=n(1001),c=(0,l.Z)(u,s,i,!1,null,null,null),d=c.exports}}]);