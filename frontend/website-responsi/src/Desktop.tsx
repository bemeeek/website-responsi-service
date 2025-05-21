import React from "react";
import headlogo1 from "./assets/headlogo.png"; // Adjust path if necessary
import "./App.css";

export const Desktop: React.FC = () => {
  return (
    <div className="desktop">
      <div className="group">
        <p className="text-wrapper text-center sm:text-left">Selamat Datang di Responsi PSI 2025</p>
        <img className="headlogo mx-auto sm:ml-0" alt="Headlogo" src={headlogo1} />
      </div>

      <div className="div text-center">Interface Tombol Ditekan</div>

      <div className="frame flex sm:grid sm:grid-cols-3 gap-4 sm:gap-6 justify-center">
        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-2">1</div>
          </div>
        </div>

        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-3">2</div>
          </div>
        </div>

        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-4">3</div>
          </div>
        </div>

        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-5">4</div>
          </div>
        </div>

        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-6">5</div>
          </div>

        </div>
        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-7">6</div>
          </div>

        </div>
        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-8">7</div>
          </div>

        </div>
        <div className="overlap-group-wrapper">
          <div className="overlap-group">
            <div className="text-wrapper-9">8</div>
          </div>
        </div>
        {/* Add other buttons here as needed */}
      </div>

      <div className="text-urutan-menjawab mt-6">Urutan Menjawab</div>

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
