"use client"
import Image from "next/image";
import { useMemo } from "react";
import dynamic from "next/dynamic";
import { useState } from "react";

interface Address {
  ISO3166_2_lvl4: string;
  city: string;
  country: string;
  country_code: string;
  county: string;
  house_number: string;
  municipality: string;
  postcode: string;
  region: string;
  road: string;
  state: string;
  state_district: string;
  suburb: string;
}

export default function Alert() {
  const Map = useMemo(() => dynamic(
    () => import('@/components/Map'),
    { 
      loading: () => <p>A map is loading</p>,
      ssr: false
    }
  ), []);

  const [address, setAddress] = useState<Address>({ 
    ISO3166_2_lvl4: '',
    city: '',
    country: '',
    country_code: '',
    county: '',
    house_number: '',
    municipality: '',
    postcode: '',
    region: '',
    road: '',
    state: '',
    state_district: '',
    suburb: ''
  });

  const updateAddress = (newAddress: Address) => {
    setAddress(newAddress);
  };

  return (
    <>
    <div className="h-[100vh] flex items-center justify-center bg-[#ededed]">
    <div className="flex justify-center items-center gap-4 w-[70%] p-4 border rounded-md shadow-lg bg-white">
    <Map address={address} updateAddress={updateAddress} />
    <div className="w-[50%] border py-12 px-6 rounded-lg">
      <div className="flex flex-col justify-center items-center mb-16">
        <p className="text-3xl font-bold text-[#093a56]">Criar um novo alerta</p>
        <span className="text-sm text-center">Para registrar um novo alerta, clique no mapa ao lado para definir sua localização atual e, em seguida, escolha o ponto exato onde deseja criar o alerta.</span>
      </div>
      <div className="flex items-center gap-1 mb-4">
        <label className="font-bold">Local de alerta selecionado:</label>
        <p className="text-sm">{address.road}, {address.suburb}, {address.city}, {address.state}, {address.country} </p>
      </div>
      <div className="flex gap-2 mb-4">
        <label className="font-bold ">Selecione o de alerta:</label>
        <select className="border rounded-md">
          <option value="1">Incêndio</option>
          <option value="2">Deslizamento</option>
          <option value="3">Enchente</option>
          <option value="4">Desabamento</option>
        </select>
      </div>
      <div className="flex items-start gap-2">
        <label className="font-bold">Descrição:</label>
        <textarea placeholder="Descreva um pouco do que foi visto..." className="text-xs border w-72 h-24 rounded-md p-1"></textarea>
      </div>
      <div className="flex justify-center mt-6">
        <button className="bg-[#093a56] text-white px-12 py-1 rounded-md hover:scale-95 transition">Enviar</button>
      </div>
    </div>
    </div>
    </div>
    </>
  );
}
