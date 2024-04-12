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

export default function NewSensor() {
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
        <p className="text-3xl font-bold text-[#093a56]">Registrar novo sensor</p>
        <span className="text-sm text-center">Para registrar um novo sensor clique no mapa e escolha ponto exato onde deseja registra-lo.</span>
      </div>
      <div className="flex items-center gap-1 mb-4">
        <label className="font-bold">Local do sensor:</label>
        <p className="text-sm">{address.road}, {address.suburb}, {address.city}, {address.state}, {address.country} </p>
      </div>
      <div className="flex items-start gap-2 mb-4">
            <label className="font-bold">Nome do sensor:</label>
            <input type="text" className="border rounded-md px-2" />
        </div>
      <div className="flex gap-2 mb-4">
        <label className="font-bold ">Tipo se sensor:</label>
        <select className="border rounded-md">
          <option value="1">MICS-6814</option>
          <option value="2">RXWLIB900</option>
        </select>
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
