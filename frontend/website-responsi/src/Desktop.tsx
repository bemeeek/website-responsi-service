import React, { useState, useEffect } from "react";
import headlogo1 from "./assets/headlogo.png"; // Adjust path if necessary
import "./App.css";

export const Desktop: React.FC = () => {
  const [buttonData, setButtonData] = useState<any>(null);
  const [loading, setLoading] = useState<boolean>(true); // To manage loading state

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

  useEffect(() => {
    const fetchButtonData = async () => {
      try {
        setLoading(true);
        const response = await fetch("http://localhost:3000/responsi");
        if (!response.ok) throw new Error("Network response was not ok");

        const data = await response.json();
        
        // Check if the data contains the expected structure
        if (data && Object.keys(data).length >= 9) {
          setButtonData(data);
        } else {
          console.warn("Unexpected data format:", data);
          setButtonData(defaultButtonData);
        }
      } catch (error) {
        console.error("Error fetching button data:", error);
        setButtonData(defaultButtonData);
      } finally {
        setLoading(false);
      }
    };

    fetchButtonData();
    const interval = setInterval(fetchButtonData, 1000); // Fetch every 1 second
    return () => clearInterval(interval);
  }, []);

  const getButtonColor = (value: number) => {
    switch (value) {
      case 1:
        return "bg-green-500"; // Button pressed (green)
      case 2:
        return "bg-yellow-500"; // First press (yellow)
      case 3:
        return "bg-blue-500"; // Second press (blue)
      default:
        return "bg-gray-500"; // Default (gray)
    }
  };

  return (
    <div className="desktop flex-col items-center justify-between h-screen p-3">
      <div className="group text-center mb-6">
        <p className="text-wrapper text-center sm:text-left">Selamat Datang di Responsi PSI 2025</p>
        <img className="headlogo mx-auto sm:ml-0" alt="Headlogo" src={headlogo1} />
      </div>

      <div className="div text-center">Interface Tombol Ditekan</div>

      {/* Loading state */}
      {loading ? (
        <div className="text-center">Loading...</div>
      ) : (
        <div className="frame">
          {buttonData ? (
            Object.entries(buttonData)
              .filter(([key]) => key.startsWith("X")) // Only render X1 to X9
              .map(([key, value], index) => (
                <div
                  key={key}
                  className={`button ${getButtonColor(Number(value))} rounded-lg flex items-center justify-center`}
                >
                  <div className="text-button">{index + 1}</div>
                </div>
              ))
          ) : (
            Object.keys(defaultButtonData).map((key, index) => (
              <div
                key={key}
                className={`button ${getButtonColor(defaultButtonData[key as keyof typeof defaultButtonData])} rounded-lg p-4 flex items-center justify-center`}
              >
                <div className="text-button">{index + 1}</div>
              </div>
            ))
          )}
        </div>
      )}

      <div className="div text-center">Urutan Menjawab</div>
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
