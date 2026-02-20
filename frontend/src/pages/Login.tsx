function Login() {
    return (
        <>
            <div className="h-screen w-screen flex justify-center p-10 bg-blue-200">
                <div className="h-fit bg-white p-5 rounded-sm">
                    <h1 className="text-center m-2">Login</h1>
                    <div className="flex flex-col gap-5">
                        <div className="flex flex-col">
                            <label htmlFor="">Email</label>
                            <input className="border" type="text" />
                        </div >
                        <div className="flex flex-col">
                            <label htmlFor="">Password</label>
                            <input className="border" type="password" />
                        </div>

                        <div className="flex flex-col">
                            <button className="bg-blue-200 rounded-lg" type="submit">Iniciar Sessão</button>
                        </div>

                    </div>
                </div>
            </div >
        </>
    )
}

export default Login;