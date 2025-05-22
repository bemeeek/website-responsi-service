import React, { useState, useEffect } from "react";
import headlogo1 from "./assets/headlogo.png";

interface ButtonData {
  X1?: number;
  Y1?: number;
  X2?: number;
  Y2?: number;
  X3?: number;
  Y3?: number;
  X4?: number;
  Y4?: number;
  X5?: number;
  Y5?: number;
  X6?: number;
  Y6?: number;
  X7?: number;
  Y7?: number;
  X8?: number;
  Y8?: number;
  X9?: number;
  Y9?: number;
}

export const Desktop: React.FC = () => {
  const [buttonData, setButtonData] = useState<ButtonData | null>(null);
  const [connectionStatus, setConnectionStatus] = useState("Menghubungkan...");

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8081/ws");

    ws.onopen = () => {
      setConnectionStatus("Terhubung");
      console.log("WebSocket connected");
    };

    ws.onmessage = (event) => {
      try {
        const data: ButtonData = JSON.parse(event.data);
        console.log("Data diterima:", data);
        setButtonData(data);
        setConnectionStatus("Data diterima");
      } catch (error) {
        console.error("Error parsing data:", error);
        setConnectionStatus("Error parsing data");
      }
    };

    ws.onerror = () => setConnectionStatus("Error koneksi");
    ws.onclose = () => setConnectionStatus("Terputus");

    return () => ws.close();
  }, []);

  const formatMilliseconds = (ms?: number) => {
    return ms ? `${ms} ms` : "-";
  };

  const getButtonColor = (buttonNumber: number) => {
    if (!buttonData) return "bg-gray-300";
    if (buttonData.X1 === buttonNumber) return "bg-green-500 animate-pulse";
    if (buttonData.X2 === buttonNumber) return "bg-yellow-500 animate-pulse";
    if (buttonData.X3 === buttonNumber) return "bg-blue-500 animate-pulse";
    return "bg-gray-300 hover:bg-gray-400 transition-colors";
  };

  // Fungsi untuk membagi array menjadi kelompok 3 elemen
  const chunkArray = (arr: number[], size: number) => {
    return Array.from({ length: Math.ceil(arr.length / size) }, (_, i) =>
      arr.slice(i * size, i * size + size)
    );
  };

  return (
    <div className="min-h-screen bg-white flex flex-col items-center p-10 space-y-8">
      {/* Header */}
      <header className="text-center space-y-4">
        <h1 className="text-3xl font-bold text-gray-800">
          Selamat Datang di Responsi PSI 2025
        </h1>
        <img 
          src={headlogo1} 
          alt="Logo" 
          className="mx-auto w-40 h-40 object-contain"
        />
        <div className="text-sm text-gray-500">
          Status: {connectionStatus}
        </div>
      </header>

      {/* Tombol Grid */}
            <section className="space-y-4 w-full max-w-4xl">
        <h2 className="text-2xl font-semibold text-center text-gray-700">
          Interface Tombol Ditekan
        </h2>
        
        <div className="grid grid-cols-3 md:grid-cols-9 gap-4 p-4 bg-gray-50 rounded-xl shadow-sm">
          {Array.from({ length: 9 }, (_, i) => {
            const buttonNumber = i + 1;
            return (
              <div 
                key={buttonNumber}
                className="flex flex-col items-center gap-2"
              >
                {/* Tombol */}
                <div className={`aspect-square w-full rounded-lg flex items-center justify-center text-4xl font-bold ${getButtonColor(buttonNumber)}`}>
                  {buttonNumber}
                </div>
                
                {/* Waktu Penekanan */}
                <div className="text-center w-full">
                  <div className="text-xs text-gray-500">Waktu</div>
                  <div className="font-mono text-blue-600 text-sm">
                    {formatMilliseconds(buttonData?.[`Y${buttonNumber}` as keyof ButtonData] as number)}
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </section>

      {/* Indikator Urutan */}
      <section className="w-full max-w-2xl space-y-4">
        <h2 className="text-2xl font-semibold text-center text-gray-700">
          Urutan Menjawab
        </h2>
        
        <div className="flex justify-center gap-8">
          {[1, 2, 3].map((order) => (
            <div key={order} className="text-center space-y-2">
              <div className={`w-20 h-20 rounded-lg ${
                (buttonData?.X1 && order === 1) ? 'bg-green-500' :
                (buttonData?.X2 && order === 2) ? 'bg-yellow-500' :
                (buttonData?.X3 && order === 3) ? 'bg-blue-500' : 'bg-gray-200'
              }`} />
              <span className="text-gray-600">Posisi {order}</span>
            </div>
          ))}
        </div>
      </section>

      {/* Footer */}
      <footer className="mt-8 text-center text-gray-500 text-sm">
        <p>
          Maafin ya ges websitenya seadanya,<br />
          yang penting fungsi bukan gengsi #anjayyyyy
        </p>
      </footer>
    </div>
  );
};