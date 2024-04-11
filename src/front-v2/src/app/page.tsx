"use client"
import Image from "next/image";
import Decotarion from "../../public/sao_paulo_decoration.svg";



export default function Home() {
  return (
    <>
    <div className="h-[100vh] flex items-center justify-center bg-[#ededed] gap-8">
    <div className="flex flex-col justify-center items-center gap-4 w-[40%] px-12 py-24 border rounded-md shadow-lg bg-white">
      <p className="font-bold text-2xl mb-8">Bem vindo ao sistema Smartopia</p>
      <div className="flex items-start gap-2">
        <label className="font-bold text-xl">Login:</label>
        <input type="text" className="border rounded-md px-2 w-full" />
      </div>
      <div className="flex items-start gap-2">
        <label className="font-bold text-xl">Senha:</label>
        <input type="password" className="border rounded-md px-2 w-full" />
      </div>
      <div className="flex justify-center mt-6">
        <button className="bg-[#093a56] text-white px-20 py-1 rounded-md hover:scale-95 transition font-bold">Entrar</button>
      </div>
    </div>
    <Image src={Decotarion} alt="São Paulo" height={370} />
    </div>
    </>
  );
}
