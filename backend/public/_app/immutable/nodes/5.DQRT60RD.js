import{s as Y,a as Z,n as X,r as ee}from"../chunks/scheduler.BfJEPAwo.js";import{S as te,i as ae,e as a,s as C,c as s,a as p,d as f,b as k,g as B,f as e,h as se,j as t,k as N,l as G}from"../chunks/index.DlzQYM85.js";import{k as le}from"../chunks/index.BoqvxspT.js";function re(I){let n,r,l,V,m,o,v,O,S,i,h,g,F="Email",U,u,q,x,_,J="Password",A,d,H,b,w,K="Sign In",$,y,Q='<a href="/recovery" class="hover:text-mango ease-in-out duration-200 inter font-normal text-neutral-500">recover account</a> <p class="inter font-normal text-neutral-300">/</p> <a href="/login" class="hover:text-mango ease-in-out duration-200 inter font-normal text-neutral-500">sign in instead</a>',z,R;return{c(){n=a("div"),r=a("div"),l=a("div"),V=C(),m=a("div"),o=a("div"),v=a("img"),S=C(),i=a("div"),h=a("div"),g=a("label"),g.textContent=F,U=C(),u=a("input"),q=C(),x=a("div"),_=a("label"),_.textContent=J,A=C(),d=a("input"),H=C(),b=a("div"),w=a("button"),w.textContent=K,$=C(),y=a("div"),y.innerHTML=Q,this.h()},l(c){n=s(c,"DIV",{class:!0});var E=p(n);r=s(E,"DIV",{class:!0});var j=p(r);l=s(j,"DIV",{class:!0}),p(l).forEach(f),V=k(j),m=s(j,"DIV",{class:!0});var W=p(m);o=s(W,"DIV",{class:!0});var L=p(o);v=s(L,"IMG",{class:!0,src:!0,alt:!0}),S=k(L),i=s(L,"DIV",{class:!0});var D=p(i);h=s(D,"DIV",{class:!0});var T=p(h);g=s(T,"LABEL",{class:!0,for:!0,"data-svelte-h":!0}),B(g)!=="svelte-ef6rif"&&(g.textContent=F),U=k(T),u=s(T,"INPUT",{class:!0,type:!0,name:!0}),T.forEach(f),q=k(D),x=s(D,"DIV",{class:!0});var M=p(x);_=s(M,"LABEL",{class:!0,for:!0,"data-svelte-h":!0}),B(_)!=="svelte-13fno33"&&(_.textContent=J),A=k(M),d=s(M,"INPUT",{class:!0,type:!0,name:!0}),M.forEach(f),H=k(D),b=s(D,"DIV",{class:!0});var P=p(b);w=s(P,"BUTTON",{class:!0,"data-svelte-h":!0}),B(w)!=="svelte-1ib6js9"&&(w.textContent=K),$=k(P),y=s(P,"DIV",{class:!0,"data-svelte-h":!0}),B(y)!=="svelte-1o89mgf"&&(y.innerHTML=Q),P.forEach(f),D.forEach(f),L.forEach(f),W.forEach(f),j.forEach(f),E.forEach(f),this.h()},h(){e(l,"class","absolute -inset-2 rounded-lg bg-gradient-to-tr from-mango to-go blur-xl"),e(v,"class","h-10 m-4"),Z(v.src,O="/mango.png")||e(v,"src",O),e(v,"alt","Mango"),e(g,"class","inter text-neutral-600"),e(g,"for","email"),e(u,"class","inter focus:shadow-black/15 ease-in-out duration-300 text-neutral-600 w-full px-2 py-1 shadow shadow-black/0 bg-neutral-200/50 outline-none rounded-md border border-neutral-400"),e(u,"type","email"),e(u,"name","email"),e(h,"class","flex flex-col gap-1 w-[75%]"),e(_,"class","inter text-neutral-600"),e(_,"for","password"),e(d,"class","inter focus:shadow-black/15 ease-in-out duration-300 text-neutral-600 w-full px-2 py-1 shadow shadow-black/0 bg-neutral-200/50 outline-none rounded-md border border-neutral-400"),e(d,"type","password"),e(d,"name","password"),e(x,"class","flex flex-col gap-1 w-[75%]"),e(w,"class","hover:bg-mango/90 ease-in-out duration-300 w-full bg-mango p-2 text-white inter font-medium rounded-full border-2 border-black/5"),e(y,"class","flex flex-row items-center justify-between"),e(b,"class","flex flex-col gap-2 w-[75%]"),e(i,"class","flex flex-col items-center gap-8 w-full"),e(o,"class","relative flex flex-col items-center py-8 shadow-inner shadow-black/25 bg-neutral-100 w-full h-full rounded-3xl p-4 scale-[99%]"),e(m,"class","relative flex size-[400px] bg-gradient-to-tr from-mango to-go rounded-3xl"),e(r,"class","relative"),e(n,"class","flex items-center justify-center h-screen")},m(c,E){se(c,n,E),t(n,r),t(r,l),t(r,V),t(r,m),t(m,o),t(o,v),t(o,S),t(o,i),t(i,h),t(h,g),t(h,U),t(h,u),N(u,I[0].email),t(i,q),t(i,x),t(x,_),t(x,A),t(x,d),N(d,I[0].password),t(i,H),t(i,b),t(b,w),t(b,$),t(b,y),z||(R=[G(u,"input",I[2]),G(d,"input",I[3]),G(w,"click",I[1])],z=!0)},p(c,[E]){E&1&&u.value!==c[0].email&&N(u,c[0].email),E&1&&d.value!==c[0].password&&N(d,c[0].password)},i:X,o:X,d(c){c&&f(n),z=!1,ee(R)}}}function ne(I,n,r){let l={email:"",password:""};async function V(){const v=await le.post("http://localhost:3000/auth/register",{json:l}).json();console.log(v)}function m(){l.email=this.value,r(0,l)}function o(){l.password=this.value,r(0,l)}return[l,V,m,o]}class de extends te{constructor(n){super(),ae(this,n,ne,re,Y,{})}}export{de as component};
