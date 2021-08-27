(self.webpackChunkdoc_ops=self.webpackChunkdoc_ops||[]).push([[4259],{3905:function(e,t,r){"use strict";r.d(t,{Zo:function(){return p},kt:function(){return f}});var n=r(7294);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function a(e,t){if(null==e)return{};var r,n,c=function(e,t){if(null==e)return{};var r,n,c={},o=Object.keys(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||(c[r]=e[r]);return c}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(n=0;n<o.length;n++)r=o[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(c[r]=e[r])}return c}var u=n.createContext({}),s=function(e){var t=n.useContext(u),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},p=function(e){var t=s(e.components);return n.createElement(u.Provider,{value:t},e.children)},l={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},d=n.forwardRef((function(e,t){var r=e.components,c=e.mdxType,o=e.originalType,u=e.parentName,p=a(e,["components","mdxType","originalType","parentName"]),d=s(r),f=c,h=d["".concat(u,".").concat(f)]||d[f]||l[f]||o;return r?n.createElement(h,i(i({ref:t},p),{},{components:r})):n.createElement(h,i({ref:t},p))}));function f(e,t){var r=arguments,c=t&&t.mdxType;if("string"==typeof e||c){var o=r.length,i=new Array(o);i[0]=d;var a={};for(var u in t)hasOwnProperty.call(t,u)&&(a[u]=t[u]);a.originalType=e,a.mdxType="string"==typeof e?e:c,i[1]=a;for(var s=2;s<o;s++)i[s]=r[s];return n.createElement.apply(null,i)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},4932:function(e,t,r){"use strict";r.r(t),r.d(t,{frontMatter:function(){return a},contentTitle:function(){return u},metadata:function(){return s},toc:function(){return p},default:function(){return d}});var n=r(2122),c=r(9756),o=(r(7294),r(3905)),i=["components"],a={},u="ISCP Architecture",s={unversionedId:"guide/core_concepts/iscp-architecture",id:"guide/core_concepts/iscp-architecture",isDocsHomePage:!1,title:"ISCP Architecture",description:"With ISCP anyone can start a own chain and define the validators. Each chain has its own state",source:"@site/docs/guide/core_concepts/iscp-architecture.md",sourceDirName:"guide/core_concepts",slug:"/guide/core_concepts/iscp-architecture",permalink:"/docs/guide/core_concepts/iscp-architecture",editUrl:"https://github.com/iotaledger/chronicle.rs/tree/main/docs/docs/guide/core_concepts/iscp-architecture.md",version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"ISCP",permalink:"/docs/guide/core_concepts/iscp"},next:{title:"Validators",permalink:"/docs/guide/core_concepts/validators"}},p=[{value:"TODO: Insert image of architecture (probably the image of page 7 of the architecture doc)",id:"todo-insert-image-of-architecture-probably-the-image-of-page-7-of-the-architecture-doc",children:[]}],l={toc:p};function d(e){var t=e.components,r=(0,c.Z)(e,i);return(0,o.kt)("wrapper",(0,n.Z)({},l,r,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"iscp-architecture"},"ISCP Architecture"),(0,o.kt)("p",null,"With ISCP anyone can start a own chain and define the validators. Each chain has its own state\nwhere a state update (going from one block to the next) is hashed and published to the main tangle\nas well by moving a special state anchor on Layer 1. This makes ISCP a more complex implementation\nof smart contracts over say Ethereum as illustrated here:"),(0,o.kt)("h2",{id:"todo-insert-image-of-architecture-probably-the-image-of-page-7-of-the-architecture-doc"},"TODO: Insert image of architecture (probably the image of page 7 of the architecture doc)"),(0,o.kt)("p",null,"A full and extensive description of the IOTA Architecture describing all components in detail can be found in this\n",(0,o.kt)("a",{parentName:"p",href:"https://github.com/iotaledger/wasp/raw/master/documentation/ISCP%20architecture%20description%20v3.pdf"},"technical description"),"."))}d.isMDXComponent=!0}}]);