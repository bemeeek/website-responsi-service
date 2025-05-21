import React, { useState, useEffect } from "react";
import headlogo1 from "./assets/headlogo.png"; // Adjust path if necessary
import "./App.css";

export const Desktop: React.FC = () => {
  // State to store the button data received from the backend
  const [buttonData, setButtonData] = useState<any>(null);

  // Default button state (this ensures buttons will still appear if no data is available)
  const defaultButtonData = {
    X1: 0,
    X2: 0,
    X3: 0,
    X4: 0,
    X5: 0,
    X6: 0,
    X7: 0,
    X8: 0,
    X9: 0,
  };

  // Fetch data from backend when the component mounts
  useEffect(() => {
    const fetchButtonData = async () => {
      try {
        const response = await fetch("http://localhost:3000/responsi"); // Adjust URL if necessary
        const data = await response.json();
        console.log("Received data:", data); // Log the data received
        setButtonData(data); // Store the button state data
      } catch (error) {
        console.error("Error fetching button data:", error);
        setButtonData(defaultButtonData); // Set default data if there's an error fetching the data
      }
    };

    fetchButtonData();
    const interval = setInterval(fetchButtonData, 1000); // Fetch data every 1 second
    return () => clearInterval(interval); // Clear the interval when the component is unmounted
  }, []);

  // Function to determine button color based on value (1 = green, 0 = gray)
  const getButtonColor = (value: number) => {
    return value === 1 ? "bg-green-500" : "bg-red-500";
  };

  return (
    <div className="desktop flex-col items-center justify-between h-screen p-3">
      <div className="group text-center mb-6">
        <p className="text-wrapper text-center sm:text-left">Selamat Datang di Responsi PSI 2025</p>
        <img className="headlogo mx-auto sm:ml-0" alt="Headlogo" src={headlogo1} />
      </div>

      <div className="div text-center">Interface Tombol Ditekan</div>
      
      {/* Frame for buttons - adjusted layout for responsiveness */}
      <div className="frame">
{buttonData ? (
  // Render tombol berdasarkan data yang diterima
  Object.entries(buttonData)
    .filter(([key]) => key.startsWith("x")) // Hanya render x1 sampai x9
    .map(([key, value], index) => (
        <div
          key={key}
          className={`button ${getButtonColor(Number(value))} rounded-lg flex items-center justify-center`}
        >
          <p className="text-black">{index + 1}</p> {/* Menampilkan angka 1, 2, 3, dst */}
        </div>
      ))
) : (
  // Fallback tombol jika tidak ada data
  Object.keys(defaultButtonData).map((key, index) => (
    <div
      key={key}
      className={`button ${getButtonColor(defaultButtonData[key as keyof typeof defaultButtonData])} rounded-lg p-4 flex items-center justify-center`}
    >
      <p className="text-white text-xl">{index + 1}</p> {/* Menampilkan angka 1, 2, 3, dst */}
    </div>
  ))
)}
      </div>
      
      <div className="text-urutan-menjawab mt-3">Urutan Menjawab</div>

      <div className="frame-2">
        <div className="frame-3">
          <div className="rectangle" />
          <div className="rectangle-2" />
        </div>

        <div className="frame-3">
          <div className="rectangle" />
          <div className="rectangle-2" />
        </div>

        <div className="frame-3">
          <div className="rectangle" />
          <div className="rectangle-2" />
        </div>
      </div>

      <p className="maafin-ya-ges text-center">
        Maafin ya ges websitenya seadanya, yang penting fungsi bukan gengsi
        <br />
        #anjayyyyy
      </p>
    </div>
  );
};
