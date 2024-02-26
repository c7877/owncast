(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[4440],{13882:function(e,n,t){"use strict";function a(e,n){if(n.length<e)throw TypeError(e+" argument"+(e>1?"s":"")+" required, but only "+n.length+" present")}t.d(n,{Z:function(){return a}})},93645:function(e,n,t){"use strict";t.d(n,{u:function(){return i}});var a={ceil:Math.ceil,round:Math.round,floor:Math.floor,trunc:function(e){return e<0?Math.ceil(e):Math.floor(e)}};function i(e){return e?a[e]:a.trunc}},59910:function(e,n,t){"use strict";t.d(n,{Z:function(){return s}});var a=t(19013),i=t(13882);function s(e,n){return(0,i.Z)(2,arguments),(0,a.Z)(e).getTime()-(0,a.Z)(n).getTime()}},11699:function(e,n,t){"use strict";t.d(n,{Z:function(){return r}});var a=t(59910),i=t(13882),s=t(93645);function r(e,n,t){(0,i.Z)(2,arguments);var r=(0,a.Z)(e,n)/1e3;return(0,s.u)(null==t?void 0:t.roundingMethod)(r)}},19013:function(e,n,t){"use strict";t.d(n,{Z:function(){return s}});var a=t(71002),i=t(13882);function s(e){(0,i.Z)(1,arguments);var n=Object.prototype.toString.call(e);return e instanceof Date||"object"===(0,a.Z)(e)&&"[object Date]"===n?new Date(e.getTime()):"number"==typeof e||"[object Number]"===n?new Date(e):(("string"==typeof e||"[object String]"===n)&&"undefined"!=typeof console&&(console.warn("Starting with v2.0.0-beta.1 date-fns doesn't accept strings as date arguments. Please use `parseISO` to parse strings. See: https://github.com/date-fns/date-fns/blob/master/docs/upgradeGuide.md#string-arguments"),console.warn(Error().stack)),new Date(NaN))}},7148:function(e,n,t){(window.__NEXT_P=window.__NEXT_P||[]).push(["/admin/config-notify",function(){return t(51375)}])},86088:function(e,n,t){"use strict";t.d(n,{Z:function(){return d}});var a=t(85893),i=t(67294),s=t(38376),r=t(70329),l=t(65326),o=t(53068),c=t(84443);let d=e=>{let{apiPath:n,checked:t,reversed:d=!1,configPath:u="",disabled:f=!1,fieldName:h,label:p,tip:m,useSubmit:g,onChange:x}=e,[b,v]=(0,i.useState)(null),y=null,{setFieldInConfigState:j}=(0,i.useContext)(c.a)||{},w=()=>{v(null),clearTimeout(y),y=null},k=async e=>{if(g){v((0,r.kg)(r.Jk));let t=d?!e:e;await (0,o.Si)({apiPath:n,data:{value:t},onSuccess:()=>{j({fieldName:h,value:t,path:u}),v((0,r.kg)(r.zv))},onError:e=>{v((0,r.kg)(r.Un,"There was an error: ".concat(e)))}}),y=setTimeout(w,o.sI)}x&&x(e)},C=null!==b&&b.type===r.Jk;return(0,a.jsxs)("div",{className:"formfield-container toggleswitch-container",children:[p&&(0,a.jsx)("div",{className:"label-side",children:(0,a.jsx)("span",{className:"formfield-label",children:p})}),(0,a.jsxs)("div",{className:"input-side",children:[(0,a.jsxs)("div",{className:"input-group",children:[(0,a.jsx)(s.Z,{className:"switch field-".concat(h),loading:C,onChange:k,defaultChecked:t,checked:t,checkedChildren:"ON",unCheckedChildren:"OFF",disabled:f}),(0,a.jsx)(l.E,{status:b})]}),(0,a.jsx)("p",{className:"field-tip",children:m})]})]})};d.defaultProps={apiPath:"",checked:!1,reversed:!1,configPath:"",disabled:!1,label:"",tip:"",useSubmit:!1,onChange:null}},55050:function(e,n,t){"use strict";var a=t(75664);n.Z=a.Z},49947:function(e,n,t){"use strict";var a=t(82215),i=t(63085),s=a.ZP;s.Header=a.h4,s.Footer=a.$_,s.Content=a.VY,s.Sider=i.Z,s._InternalSiderContext=i.D,n.default=s},6647:function(e,n,t){"use strict";var a=t(37525);n.Z=a.Z},38376:function(e,n,t){"use strict";t.d(n,{Z:function(){return y}});var a=t(87462),i=t(4942),s=t(38813),r=t(93967),l=t.n(r),o=t(97685),c=t(45987),d=t(67294),u=t(21640),f=t(79097),h=d.forwardRef(function(e,n){var t,a=e.prefixCls,s=void 0===a?"rc-switch":a,r=e.className,h=e.checked,p=e.defaultChecked,m=e.disabled,g=e.loadingIcon,x=e.checkedChildren,b=e.unCheckedChildren,v=e.onClick,y=e.onChange,j=e.onKeyDown,w=(0,c.Z)(e,["prefixCls","className","checked","defaultChecked","disabled","loadingIcon","checkedChildren","unCheckedChildren","onClick","onChange","onKeyDown"]),k=(0,u.Z)(!1,{value:h,defaultValue:p}),C=(0,o.Z)(k,2),N=C[0],Z=C[1];function E(e,n){var t=N;return m||(Z(t=e),null==y||y(t,n)),t}var S=l()(s,r,(t={},(0,i.Z)(t,"".concat(s,"-checked"),N),(0,i.Z)(t,"".concat(s,"-disabled"),m),t));return d.createElement("button",Object.assign({},w,{type:"button",role:"switch","aria-checked":N,disabled:m,className:S,ref:n,onKeyDown:function(e){e.which===f.Z.LEFT?E(!1,e):e.which===f.Z.RIGHT&&E(!0,e),null==j||j(e)},onClick:function(e){var n=E(!N,e);null==v||v(n,e)}}),g,d.createElement("span",{className:"".concat(s,"-inner")},N?x:b))});h.displayName="Switch";var p=t(71946),m=t(6089),g=t(41395),x=t(40823),b=function(e,n){var t={};for(var a in e)Object.prototype.hasOwnProperty.call(e,a)&&0>n.indexOf(a)&&(t[a]=e[a]);if(null!=e&&"function"==typeof Object.getOwnPropertySymbols)for(var i=0,a=Object.getOwnPropertySymbols(e);i<a.length;i++)0>n.indexOf(a[i])&&Object.prototype.propertyIsEnumerable.call(e,a[i])&&(t[a[i]]=e[a[i]]);return t},v=d.forwardRef(function(e,n){var t=e.prefixCls,r=e.size,o=e.disabled,c=e.loading,u=e.className,f=b(e,["prefixCls","size","disabled","loading","className"]),v=d.useContext(p.E_),y=v.getPrefixCls,j=v.direction,w=d.useContext(g.Z),k=d.useContext(m.Z),C=(null!=o?o:k)||c,N=y("switch",t),Z=d.createElement("div",{className:"".concat(N,"-handle")},c&&d.createElement(s.Z,{className:"".concat(N,"-loading-icon")})),E=l()((0,i.Z)((0,i.Z)((0,i.Z)({},"".concat(N,"-small"),"small"===(r||w)),"".concat(N,"-loading"),c),"".concat(N,"-rtl"),"rtl"===j),void 0===u?"":u);return d.createElement(x.Z,{insertExtraNode:!0},d.createElement(h,(0,a.Z)({},f,{prefixCls:N,className:E,disabled:C,ref:n,loadingIcon:Z})))});v.__ANT_SWITCH=!0;var y=v},51375:function(e,n,t){"use strict";t.r(n),t.d(n,{default:function(){return S}});var a=t(85893),i=t(47271),s=t(46994),r=t(6647),l=t(55050),o=t(90622),c=t(67294),d=t(41664),u=t.n(d),f=t(84443),h=t(62376),p=t(65326),m=t(53068),g=t(86088),x=t(70329);let{Title:b}=i.default,v=()=>{let{serverConfig:e,setFieldInConfigState:n}=(0,c.useContext)(f.a)||{},{notifications:t}=e||{},{discord:i}=t||{},{enabled:s,webhook:r,goLiveMessage:l}=i||{},[d,u]=(0,c.useState)({}),[v,y]=(0,c.useState)(null),[j,w]=(0,c.useState)(!1);(0,c.useEffect)(()=>{u({enabled:s,webhook:r,goLiveMessage:l})},[t,i]);let k=()=>""!==r&&""!==l,C=e=>{let{fieldName:n,value:t}=e;u({...d,[n]:t}),w(k())},N=()=>{y(null),clearTimeout(null)},Z=async()=>{await (0,m.Si)({apiPath:"/notifications/discord",data:{value:d},onSuccess:()=>{n({fieldName:"discord",value:d,path:"notifications"}),y((0,x.kg)(x.zv,"Updated.")),setTimeout(N,m.sI)},onError:e=>{y((0,x.kg)(x.Un,e)),setTimeout(N,m.sI)}})};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(b,{children:"Discord"}),(0,a.jsx)("p",{className:"description reduced-margins",children:"Let your Discord channel know each time you go live."}),(0,a.jsxs)("p",{className:"description reduced-margins",children:[(0,a.jsx)("a",{href:"https://support.discord.com/hc/en-us/articles/228383668",target:"_blank",rel:"noreferrer",children:"Create a webhook"})," ","under ",(0,a.jsx)("i",{children:"Edit Channel / Integrations"})," on your Discord channel and provide it below."]}),(0,a.jsx)(g.Z,{apiPath:"",fieldName:"discordEnabled",label:"Enable Discord",checked:d.enabled,onChange:e=>{C({fieldName:"enabled",value:e})}}),(0,a.jsx)("div",{style:{display:d.enabled?"block":"none"},children:(0,a.jsx)(h.nv,{...m.oy.webhookUrl,required:!0,value:d.webhook,onChange:C})}),(0,a.jsx)("div",{style:{display:d.enabled?"block":"none"},children:(0,a.jsx)(h.nv,{...m.oy.goLiveMessage,required:!0,value:d.goLiveMessage,onChange:C})}),(0,a.jsx)(o.Z,{type:"primary",onClick:Z,style:{display:j?"inline-block":"none",position:"relative",marginLeft:"auto",right:"0",marginTop:"20px"},children:"Save"}),(0,a.jsx)(p.E,{status:v})]})},{Title:y}=i.default,j=()=>{let{serverConfig:e,setFieldInConfigState:n}=(0,c.useContext)(f.a)||{},{notifications:t}=e||{},{browser:i}=t||{},{enabled:s,goLiveMessage:r}=i||{},[l,d]=(0,c.useState)({}),[u,b]=(0,c.useState)(null),[v,j]=(0,c.useState)(!1);(0,c.useEffect)(()=>{d({enabled:s,goLiveMessage:r})},[t,i]);let w=()=>!0,k=e=>{let{fieldName:n,value:t}=e;console.log(n,t),d({...l,[n]:t}),j(w())},C=()=>{b(null),clearTimeout(null)},N=async()=>{await (0,m.Si)({apiPath:"/notifications/browser",data:{value:l},onSuccess:()=>{n({fieldName:"browser",value:l,path:"notifications"}),b((0,x.kg)(x.zv,"Updated.")),setTimeout(C,m.sI)},onError:e=>{b((0,x.kg)(x.Un,e)),setTimeout(C,m.sI)}})};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(y,{children:"Browser Alerts"}),(0,a.jsx)("p",{className:"description reduced-margins",children:"Viewers can opt into being notified when you go live with their browser."}),(0,a.jsx)("p",{className:"description reduced-margins",children:"Not all browsers support this."}),(0,a.jsx)(g.Z,{apiPath:"",fieldName:"enabled",label:"Enable browser notifications",onChange:e=>{k({fieldName:"enabled",value:e})},checked:l.enabled}),(0,a.jsx)("div",{style:{display:l.enabled?"block":"none"},children:(0,a.jsx)(h.nv,{...m.mv.goLiveMessage,required:!0,type:h.Sk,value:l.goLiveMessage,onChange:k})}),(0,a.jsx)(o.Z,{type:"primary",style:{display:v?"inline-block":"none",position:"relative",marginLeft:"auto",right:"0",marginTop:"20px"},onClick:N,children:"Save"}),(0,a.jsx)(p.E,{status:u})]})},{Title:w}=i.default,k=()=>{let{serverConfig:e}=(0,c.useContext)(f.a)||{},{federation:n}=e||{},{enabled:t}=n||{},[i,s]=(0,c.useState)({});return(0,c.useEffect)(()=>{s({enabled:t})},[t]),(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(w,{children:"Fediverse Social"}),(0,a.jsx)("p",{className:"description",children:"Enabling the Fediverse social features will not just alert people to when you go live, but also enable other functionality."}),(0,a.jsxs)("p",{children:["Fediverse social features:"," ",(0,a.jsx)("span",{style:{color:n.enabled?"green":"red"},children:i.enabled?"Enabled":"Disabled"})]}),(0,a.jsx)(u(),{passHref:!0,href:"/admin/config-federation/",children:(0,a.jsx)(o.Z,{type:"primary",style:{position:"relative",marginLeft:"auto",right:"0",marginTop:"20px"},children:"Configure"})})]})};var C=t(55926),N=t(94956),Z=t(695);let{Title:E}=i.default;function S(){let[e,n]=(0,c.useState)(null),{serverConfig:t}=(0,c.useContext)(f.a)||{},{yp:i}=t,{instanceUrl:d}=i,[h,p]=(0,c.useState)(!1);(0,c.useEffect)(()=>{n({instanceUrl:d})},[i]);let g=""!==d,x=!g&&(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(s.Z,{message:"You must set your server URL before you can enable this feature.",type:"warning",showIcon:!0}),(0,a.jsx)("br",{}),(0,a.jsx)(C.$7,{fieldName:"instanceUrl",...m.yi,value:(null==e?void 0:e.instanceUrl)||"",initialValue:i.instanceUrl,type:C.xA,onChange:t=>{let{fieldName:a,value:i}=t;p((0,N.jv)(i)),n({...e,[a]:i})},onSubmit:()=>{h&&n({...e,enabled:!1})},required:!0})]});return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(E,{children:"Notifications"}),(0,a.jsxs)("p",{className:"description",children:["Let your viewers know when you go live by supporting any of the below notification channels."," ",(0,a.jsx)("a",{href:"https://owncast.online/docs/notifications/?source=admin",target:"_blank",rel:"noopener noreferrer",children:"Learn more about live notifications."})]}),x,(0,a.jsxs)(r.Z,{children:[(0,a.jsx)(l.Z,{span:10,className:"form-module ".concat(g?"":"disabled"),style:{margin:"5px",display:"flex",flexDirection:"column"},children:(0,a.jsx)(j,{})}),(0,a.jsx)(l.Z,{span:10,className:"form-module ".concat(g?"":"disabled"),style:{margin:"5px",display:"flex",flexDirection:"column"},children:(0,a.jsx)(v,{})}),(0,a.jsx)(l.Z,{span:10,className:"form-module ".concat(g?"":"disabled"),style:{margin:"5px",display:"flex",flexDirection:"column"},children:(0,a.jsx)(k,{})}),(0,a.jsxs)(l.Z,{span:10,className:"form-module ".concat(g?"":"disabled"),style:{margin:"5px",display:"flex",flexDirection:"column"},children:[(0,a.jsx)(E,{children:"Custom"}),(0,a.jsx)("p",{className:"description",children:"Build your own notifications by using custom webhooks."}),(0,a.jsx)(u(),{passHref:!0,href:"/admin/webhooks",children:(0,a.jsx)(o.Z,{type:"primary",style:{position:"relative",marginLeft:"auto",right:"0",marginTop:"20px"},children:"Create"})})]})]})]})}S.getLayout=function(e){return(0,a.jsx)(Z.l,{page:e})}},11163:function(e,n,t){e.exports=t(73035)}},function(e){e.O(0,[6410,8768,947,7406,4716,2862,7271,9083,811,695,2888,9774,179],function(){return e(e.s=7148)}),_N_E=e.O()}]);