(function(e){function t(t){for(var a,i,o=t[0],c=t[1],l=t[2],d=0,m=[];d<o.length;d++)i=o[d],Object.prototype.hasOwnProperty.call(s,i)&&s[i]&&m.push(s[i][0]),s[i]=0;for(a in c)Object.prototype.hasOwnProperty.call(c,a)&&(e[a]=c[a]);u&&u(t);while(m.length)m.shift()();return r.push.apply(r,l||[]),n()}function n(){for(var e,t=0;t<r.length;t++){for(var n=r[t],a=!0,o=1;o<n.length;o++){var c=n[o];0!==s[c]&&(a=!1)}a&&(r.splice(t--,1),e=i(i.s=n[0]))}return e}var a={},s={app:0},r=[];function i(t){if(a[t])return a[t].exports;var n=a[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,i),n.l=!0,n.exports}i.m=e,i.c=a,i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)i.d(n,a,function(t){return e[t]}.bind(null,a));return n},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/";var o=window["webpackJsonp"]=window["webpackJsonp"]||[],c=o.push.bind(o);o.push=t,o=o.slice();for(var l=0;l<o.length;l++)t(o[l]);var u=c;r.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"034f":function(e,t,n){"use strict";var a=n("85ec"),s=n.n(a);s.a},"43c8":function(e,t,n){},"4c32":function(e,t,n){"use strict";var a=n("43c8"),s=n.n(a);s.a},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var a=n("2b0e"),s=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("transition",{attrs:{name:"fade",mode:"out-in",appear:""}},[n("keep-alive",{attrs:{exclude:"checkout"}},[n("router-view",{key:e.$route.fullPath})],1)],1)},r=[],i={name:"App"},o=i,c=(n("034f"),n("2877")),l=Object(c["a"])(o,s,r,!1,null,null,null),u=l.exports,d=n("8c4f"),m=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("div",{staticClass:"container max-width-adaptive-md margin-bottom-lg"},[e.isLoading?e._e():n("events-list",{attrs:{events:e.computedEvents},scopedSlots:e._u([{key:"default",fn:function(t){return n("li",{staticClass:"stack-cards__item bg radius-lg shadow-md js-stack-cards__item",attrs:{"data-theme":"default"}},[n("div",{staticClass:"grid"},[n("div",{staticClass:"col-6 flex items-center height-100%"},[n("div",{staticClass:"text-component padding-md"},[n("h2",[e._v(e._s(t.name))]),n("p",{staticClass:"display@xs"},[e._v("Address: "+e._s(t.address))]),n("p",{staticClass:"display@xs"},[e._v("Event Prices:")]),e._l(t.event_kinds,(function(t){return n("div",{key:t.id},[n("ul",[n("li",[e._v(e._s(t.name)+" for ₦ "+e._s(parseFloat(t.price).toFixed(2)))])])])})),n("br"),n("p",[n("router-link",{staticClass:"btn btn--accent",attrs:{to:{name:"event",params:{code:t.code}}}},[e._v("Attend Event")])],1)],2)]),n("div",{staticClass:"col-6 height-100%"},[n("img",{staticClass:"block width-100% height-100% object-cover",attrs:{src:t.image},on:{error:e.myDefaultImage}})])])])}}],null,!1,1172296340)})],1),e._m(0)])},v=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"container max-width-adaptive-sm"},[n("div",{staticClass:"text-component line-height-lg v-space-md"},[n("p")])])}],p=(n("fb6a"),n("96cf"),n("1da1")),f=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"even--list"},e._l(e.events,(function(t){return n("p",e._b({key:t.id},"p",{event:t},!1),[e._t("default",null,null,t),n("br"),n("br")],2)})),0)},h=[],g={name:"events-list",props:{events:{type:Array,required:!0}}},_=g,b=Object(c["a"])(_,f,h,!1,null,null,null),y=b.exports,k=(n("d3b7"),n("bc3a")),w=n.n(k),x="/api",C="FLWPUBK_TEST-ef3f5d76636c8933f6f6a358ca27ca62-X";w.a.defaults.baseURL=x;var j=w.a.create({API_URL:x,headers:{}});j.interceptors.response.use(null,(function(e){var t="/error";switch(e.response.status){case 401:t="/login";break;case 404:t="/404";break}return ce.push(t),Promise.reject(e)}));var E=j,P="/events",O={get:function(){return E.get("".concat(P))},getEvent:function(e){return E.get("".concat(P,"/").concat(e))},getTicket:function(e){return E.get("".concat(P,"/type/").concat(e))},processPayment:function(e){return E.post("".concat(P,"/payment"),e)},verifyPayment:function(e){return E.post("".concat(P,"/payment/verify"),e)}},R={events:O},T={get:function(e){return R[e]}},L=T.get("events"),S={name:"Home",components:{EventsList:y},data:function(){return{isLoading:!1,events:[]}},created:function(){this.getEvents()},methods:{getEvents:function(){var e=Object(p["a"])(regeneratorRuntime.mark((function e(){var t,n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return this.isLoading=!0,e.next=3,L.get();case 3:t=e.sent,n=t.data,this.isLoading=!1,this.events=n.data;case 7:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),myDefaultImage:function(e){e.target.src="assets/img/img-1.jpg"}},computed:{computedEvents:function(){return this.events.slice(0,10)}}},D=S,I=(n("6c99"),Object(c["a"])(D,m,v,!1,null,"339ed076",null)),$=I.exports,N=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("div",{staticClass:"text-component text-center"},[n("p",[n("small",[n("router-link",{attrs:{to:"/"}},[e._v("Go home")])],1)])]),n("div",{staticClass:"container max-width-adaptive-md margin-bottom-lg"},[this.isLoading?n("div",[n("li",{staticClass:"stack-cards__item bg radius-lg shadow-md js-stack-cards__item",attrs:{"data-theme":"default"}})]):n("div",[n("li",{staticClass:"stack-cards__item bg radius-lg shadow-md js-stack-cards__item",attrs:{"data-theme":"default"}},[n("div",{staticClass:"grid"},[n("div",{staticClass:"col-6 flex items-center height-100%"},[n("div",{staticClass:"text-component padding-md "},[n("h2",[e._v(e._s(e.event.name))]),n("p",{staticClass:"display@xs"},[e._v("Address: "+e._s(e.event.address))]),n("p",{staticClass:"display@xs"},[e._v("Date: "+e._s(e.event.date))]),e.hasTicket?n("div",[n("p",{staticClass:"display@xs"},[e._v("Select ticket type ")]),n("div",[n("ul",[n("v-select",{attrs:{label:"name",options:e.event.event_kinds},on:{input:e.setSelected}})],1)]),n("br"),n("br")]):n("div",[e._v(" We're sorry. There are no ticket(s) available for this event yet. "),n("br"),n("br"),n("p",[n("router-link",{staticClass:"btn btn--accent",attrs:{to:"/"}},[e._v("Back to Events")])],1)])])]),n("div",{staticClass:"col-6 height-100%"},[n("img",{staticClass:"block lazy width-100% height-100% object-cover",attrs:{src:e.event.image},on:{error:e.myDefaultImage}})])])])])]),e._m(0)])},A=[function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{staticClass:"container max-width-adaptive-sm"},[n("div",{staticClass:"text-component line-height-lg v-space-md"},[n("p")])])}],G=(n("6dfc"),T.get("events")),F={name:"event",data:function(){return{isLoading:!1,event:[],options:[]}},created:function(){this.getEvent()},methods:{getEvent:function(){var e=Object(p["a"])(regeneratorRuntime.mark((function e(){var t,n;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return this.isLoading=!0,e.next=3,G.getEvent(this.$route.params.code.toLowerCase());case 3:t=e.sent,n=t.data,this.isLoading=!1,console.log(this.isLoading),this.event=n.data;case 8:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),myDefaultImage:function(e){e.target.src="assets/img/img-1.jpg"},setSelected:function(e){console.log(e.id),localStorage.setItem("event_id",e.id),ce.push("/checkout")}},computed:{computedEvents:function(){return this.event.slice(0,10)},hasTicket:function(){return this.event.event_kinds.length}}},U=F,B=(n("4c32"),Object(c["a"])(U,N,A,!1,null,"0b412d01",null)),M=B.exports,q=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[e.event_details.active?n("div",{staticClass:"container max-width-adaptive-md margin-bottom-lg"},[n("center",[e.isLoading?n("p",[n("center",[n("small",[e._v("Loading......")])])],1):n("p",[n("small",[n("router-link",{attrs:{to:{name:"event",params:{code:e.event_details.code}}}},[e._v("Go back")])],1)])]),e.success?n("b-alert",{attrs:{variant:"success",show:""}},[e._v("Payment was completed successfully")]):e._e(),e.failed?n("b-alert",{attrs:{variant:"warning",show:""}},[e._v(e._s(this.error?this.error:"Something went wrong with this payment, please try again or send an email to hello@tonidev.net"))]):e._e(),e.show&&!e.startPayment?n("b-form",{on:{submit:[e.onSubmit,function(e){e.stopPropagation(),e.preventDefault()}],reset:e.onReset}},[n("b-form-group",{attrs:{id:"input-group-1",label:"Email address:","label-for":"input-1",description:"We'll never share your email with anyone else."}},[n("b-form-input",{attrs:{id:"input-1",type:"email",required:"",placeholder:"Enter email"},model:{value:e.form.email,callback:function(t){e.$set(e.form,"email",t)},expression:"form.email"}})],1),n("b-form-group",{attrs:{id:"input-group-2",label:"Your Name:","label-for":"input-2"}},[n("b-form-input",{attrs:{id:"input-2",state:e.validation,required:"",placeholder:"Enter name"},model:{value:e.form.name,callback:function(t){e.$set(e.form,"name",t)},expression:"form.name"}}),n("b-form-invalid-feedback",{attrs:{state:e.validation}},[e._v("Your name must be 4-30 characters long.")]),n("b-form-valid-feedback",{attrs:{state:e.validation}},[e._v("You've got a cute name :(")])],1),n("h4",[e._v("Details of Payment")]),n("p",[e._v(" Title of Event : "+e._s(e.event_details.event_name)+" "),n("br"),e._v(" Amount: NGN "+e._s(e.event_details.price)+".00 "),n("br"),e._v(" Date: "+e._s(e.event_details.date)+" ")]),n("b-button",{attrs:{disabled:e.clicked,type:"submit",variant:"primary"}},[e._v(e._s(this.clicked?"Processsing ..... ":"Buy Ticket.."))]),n("b-button",{attrs:{type:"reset",variant:"danger"}},[e._v("Reset")])],1):e._e(),n("center",[e.startPayment?n("div",[n("h4",[e._v("Details of Payment")]),n("p",[e._v(" Title of Event : "+e._s(e.event_details.event_name)+" "),n("br"),e._v(" Amount: NGN "+e._s(e.event_details.price)+".00 "),n("br"),e._v(" Date: "+e._s(e.event_details.date)+" ")]),e.paymentComplete?n("div",[n("a",{staticClass:"downloadButton",attrs:{href:e.buttonUrl,target:"blank"}},[e._v(" Download Receipt/Ticket ")])]):n("div",[n("Rave",{attrs:{email:e.payment.email,amount:e.payment.amount,reference:e.payment.reference,"rave-key":e.raveKey,callback:e.callback,close:e.close,metadata:e.meta,customerFirstname:e.form.name,paymentOptions:"card,barter,account,ussd",hostedPayemt:"1",customTitle:"Golang Test",currency:"NGN",country:"NG"}})],1)]):e._e()])],1):n("small",[n("center",[n("router-link",{attrs:{to:"/"}},[e._v("Go home")]),n("br"),e._v("This event does not exist or this ticket kind is sold out "),n("br")],1)],1)])},J=[],K=(n("b0c0"),n("ce91")),W=n.n(K),Y=T.get("events"),z=C,H={name:"checkout",components:{Rave:W.a},data:function(){return{raveKey:z,isLoading:!1,failed:!1,buttonUrl:"",error:"",success:!1,paymentComplete:!1,event_details:[],payment:[],meta:[{ticket:parseInt(localStorage.getItem("event_id"))}],form:{email:"",name:"",event_id:parseInt(localStorage.getItem("event_id"))},button:{text:"Buy Ticket"},clicked:!1,show:!0,startPayment:!1}},created:function(){this.getTicket()},methods:{getTicket:function(){var e=Object(p["a"])(regeneratorRuntime.mark((function e(){var t,n,a;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:if(this.isLoading=!0,t=localStorage.getItem("event_id"),!t||!parseInt(t,10)){e.next=9;break}return e.next=5,Y.getTicket(t);case 5:n=e.sent,a=n.data,this.isLoading=!1,this.event_details=a.data;case 9:case"end":return e.stop()}}),e,this)})));function t(){return e.apply(this,arguments)}return t}(),onSubmit:function(){var e=Object(p["a"])(regeneratorRuntime.mark((function e(t){var n,a,s;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return t.preventDefault(),this.clicked=!this.clicked,n=JSON.stringify(this.form),e.next=5,Y.processPayment(n);case 5:if(a=e.sent,s=a.data,this.failed=!1,200!=!s.status){e.next=13;break}return this.failed=!0,this.error=s.message,this.clicked=!1,e.abrupt("return");case 13:this.payment=s.data,this.clicked=!1,this.startPayment=!0;case 16:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),callback:function(){var e=Object(p["a"])(regeneratorRuntime.mark((function e(t){var n,a,s;return regeneratorRuntime.wrap((function(e){while(1)switch(e.prev=e.next){case 0:return n=t,a=t.tx.txRef,e.next=4,Y.verifyPayment({txref:a});case 4:s=e.sent,200==!s.data.status?this.failed=!0:("00"==n.tx.chargeResponseCode||"0"==n.tx.chargeResponseCode?(this.paymentComplete=!0,this.buttonUrl=s.data.data,this.success=!0,window.close()):(this.failed=!0,console.log("Faileddd!!!")),window.close()),window.close();case 7:case"end":return e.stop()}}),e,this)})));function t(t){return e.apply(this,arguments)}return t}(),close:function(){console.log("Payment closed")},onReset:function(e){var t=this;e.preventDefault(),this.form.email="",this.form.name="",this.show=!1,this.$nextTick((function(){t.show=!0}))},myDefaultImage:function(e){e.target.src="assets/img/img-1.jpg"}},computed:{computedEvents:function(){return this.event_details},validation:function(){return this.form.name.length>4&&this.form.name.length<30}}},X=H,Q=(n("d1a8"),Object(c["a"])(X,q,J,!1,null,null,null)),V=Q.exports,Z=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[e._v(" 404 ")])},ee=[],te={name:"PageNotFound"},ne=te,ae=Object(c["a"])(ne,Z,ee,!1,null,null,null),se=ae.exports,re=n("4a7a"),ie=n.n(re),oe=n("5f5b");n("f9e3"),n("2dd8");a["default"].component("v-select",ie.a),a["default"].use(d["a"]),a["default"].use(oe["a"]);var ce=new d["a"]({routes:[{path:"/",name:"home",component:$},{path:"/event/:code",name:"event",component:M},{path:"/checkout",name:"checkout",component:V},{path:"*",component:se}]});a["default"].config.productionTip=!1,new a["default"]({el:"#app",router:ce,render:function(e){return e(u)}})},"6c99":function(e,t,n){"use strict";var a=n("f544"),s=n.n(a);s.a},"85ec":function(e,t,n){},d1a8:function(e,t,n){"use strict";var a=n("ebbd"),s=n.n(a);s.a},ebbd:function(e,t,n){},f544:function(e,t,n){}});
//# sourceMappingURL=app.8e39fecb.js.map