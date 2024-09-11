<script>
    import ky from "ky";

    let user = {
        email: "",
        password: "",
    }

    async function register() {
        const res = await ky.post("http://localhost:3000/auth/register", {
            json: user
        }).json();

        console.log(res);
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
                        <input class="inter focus:shadow-black/15 ease-in-out duration-300 text-neutral-600 w-full px-2 py-1 shadow shadow-black/0 bg-neutral-200/50 outline-none rounded-md border border-neutral-400" bind:value={user.email} type="email" name="email">
                    </div>
                    <div class="flex flex-col gap-1 w-[75%]">
                        <label class="inter text-neutral-600" for="password">Password</label>
                        <input class="inter focus:shadow-black/15 ease-in-out duration-300 text-neutral-600 w-full px-2 py-1 shadow shadow-black/0 bg-neutral-200/50 outline-none rounded-md border border-neutral-400" bind:value={user.password} type="password" name="password">
                    </div>
                    <div class="flex flex-col gap-2 w-[75%]">
                        <button on:click={register} class="hover:bg-mango/90 ease-in-out duration-300 w-full bg-mango p-2 text-white inter font-medium rounded-full border-2 border-black/5">Sign In</button>
                        <div class="flex flex-row items-center justify-between">
                            <a href="/recovery" class="hover:text-mango ease-in-out duration-200 inter font-normal text-neutral-500">recover account</a>
                            <p class="inter font-normal text-neutral-300">/</p>
                            <a href="/login" class="hover:text-mango ease-in-out duration-200 inter font-normal text-neutral-500">sign in instead</a>
                        </div>
                    </div>
                </div>
              </div>
        </div>
    </div>
</div>