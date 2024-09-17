<script>
    import ky from "ky";
    import { toast } from "svelte-sonner";
    import Eye from "lucide-svelte/icons/eye";

    let showPassword = false;
    let isRequestingLogin = false;
    let user = {
        email: "",
        password: "",
    }

    async function login() {
        isRequestingLogin = true;
        
        const res = await ky.post("http://localhost:3000/auth/login", {
            json: user
        }).json();

        if (res.code == 200) toast.success(res.message);
        else toast.error(res.message);

        isRequestingLogin = false;
    }
</script>

<div class="flex items-center justify-center h-screen">
    <div class="relative">
        <div class="absolute -inset-2 rounded-lg bg-gradient-to-tr from-mango to-go blur-xl"></div>
        <div class="relative flex size-[400px] bg-gradient-to-tr from-mango to-go rounded-3xl">
            <div class="relative flex flex-col items-center py-8 shadow-inner shadow-black/25 bg-neutral-100 w-full h-full rounded-3xl p-4 scale-[99%]">
                <img class="h-10 m-4" src="/mango.png" alt="Mango">
                <div class="flex flex-col items-center gap-8 w-full">
                    <div class="flex flex-col gap-1 w-[75%]">
                        <label class="inter text-neutral-600" for="email">Email</label>
                        <input class="shadow shadow-black/5 ease-in-out duration-300 text-neutral-600 w-full px-2 py-1 bg-neutral-200/50 outline-none rounded-md border border-neutral-400" bind:value={user.email} type="email" name="email">
                    </div>
                    <div class="flex flex-col gap-1 w-[75%]">
                        <label class="inter text-neutral-600" for="password">Password</label>
                        <div class="shadow shadow-black/5 text-neutral-600 w-full flex flex-row items-center bg-neutral-200/50 border border-neutral-400 px-2 py-1 rounded-md">
                            <input class="inter text-neutral-600 bg-transparent w-full outline-none" type={showPassword ? "text" : "password"} name="password">
                            <button 
                                on:mousedown={() => showPassword = true}
                                on:mouseup={() => showPassword = false}
                                on:mouseleave={() => showPassword = false}
                            >
                                <Eye class="hover:text-neutral-600 ease-in-out duration-200 cursor-pointer text-neutral-500" />
                            </button>
                        </div>
                    </div>
                    <div class="flex flex-col gap-2 w-[75%]">
                        {#if isRequestingLogin}
                            <button on:click={login} disabled class="hover:bg-mango/90 pulse ease-in-out duration-300 w-full bg-mango p-2 text-white inter font-medium rounded-full border-2 border-black/5">Sign In</button>
                        {:else}
                            <button on:click={login} class="hover:bg-mango/90 ease-in-out duration-300 w-full bg-mango p-2 text-white inter font-medium rounded-full border-2 border-black/5">Sign In</button>
                        {/if}
                        <div class="flex flex-row items-center justify-between">
                            <a href="/recovery" class="hover:text-mango ease-in-out duration-200 inter font-normal text-neutral-500">recover account</a>
                            <p class="inter font-normal text-neutral-300">/</p>
                            <a href="/register" class="hover:text-mango ease-in-out duration-200 inter font-normal text-neutral-500">sign up instead</a>
                        </div>
                    </div>
                </div>
              </div>
        </div>
    </div>
</div>

<style>
    @keyframes pulse {
        50% {
            opacity: .5;
        }
    }
    .pulse {
        animation: pulse 1s cubic-bezier(0.4, 0, 0.6, 1) infinite;
    }
</style>