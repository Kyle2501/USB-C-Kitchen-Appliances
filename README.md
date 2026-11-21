# USB-C-Kitchen-Appliances
.
https://github.com/Kyle2501/USB-C-Kitchen-Appliances

--- 
_ `...` ~

+ ice maker
+ espresso machine 
+ hot water boiler
+ air dryer / dehydrator 
+ vegetable slicer 
+ cheese grater 
+ pressure cooker 
+ rice steamer 
+ room fan
+ count down timer

____

# 🍲🎭 USB-C Kitchen Appliances
**Antistatic Lab × Ejected Media / localSq Research**

> Modular, low-voltage culinary hardware & firmware specs powered by USB-C PD 3.1 (EPR / PPS) — built for precision warming, compact prep actuation, and safe thermal telemetry.

---

## 🎯 Quick Overview (TL;DR)
* **Goal:** Replace bulky AC kitchen cords with safe, portable, open-spec USB-C PD 3.1 appliances.
* **Power Standard:** USB-C PD 3.0 / 3.1 EPR (up to $28\text{ V} - 48\text{ V}$, $100\text{ W} - 240\text{ W}$) & PPS programmable rails.
* **Design Philosophy:** AuDHD-friendly modular isolation — every hardware subsystem separates **Power**, **Heat**, **Actuation**, and **Safety Telemetry**.

---

## 🗂️ Repository Architecture

```text
USB-C-Kitchen-Appliances/
├── 01_pd_power_rail/
│   ├── usb_c_epr_28v_48v.hw      # High-voltage PD 3.1 sink & bus protection
│   ├── pps_sink_negotiator.ic    # Dynamic voltage/current contract controller
│   └── thermal_cutoff_guard.io   # Hardware brownout & over-current lockout
├── 02_induction_heat/
│   ├── low_watt_simmer_plate.hw  # Flat-coil PID simmer & warming surface
│   ├── sous_vide_stick.hw        # Immersion heater & impeller circulation
│   └── brie_melt_warmer.hw       # Gentle 55°C–65°C cheese melt cradle
├── 03_prep_actuation/
│   ├── precision_spice_mill.hw   # High-torque brushless micro-grinder
│   ├── herb_garlic_chopper.hw    # Interlocked pulse-blade prep module
│   └── dispensing_valve.io       # Food-grade peristaltic / pinch valve
└── 04_telemetry_safety/
    ├── ntc_probe_array.io        # Dual-redundant 100k NTC thermal sensing
    ├── current_sense_shunt.ic    # Real-time wattage & stall detection
    └── auto_shutoff_watchdog.io  # Independent hardware timer & dry-boil trip
```

---

## ⚡ Module Breakdown

### 01 · PD Power & Rail (`01_pd_power_rail/`)
* **What it does:** Negotiates clean DC power from any compliant USB-C power bank or wall GaN brick.
* **Key Specs:**
  * Supports $5\text{ V}$, $9\text{ V}$, $15\text{ V}$, $20\text{ V}$, $28\text{ V}$, $36\text{ V}$, and $48\text{ V}$ PDOs.
  * PPS stepping in $20\text{ mV}$ increments for direct resistive heater throttling without heavy switching noise.
* **Safety Guard:** Automatic soft-start and brownout disconnect if source voltage dips $>10\%$.

### 02 · Induction & Heat (`02_induction_heat/`)
* **What it does:** Converts negotiated DC power into controlled, scorch-free culinary heat.
* **Appliance Profiles:**
  * `brie_melt_warmer.hw`: Holds Marin French Brie wheels at an ideal melt window ($58^\circ\text{C} \pm 1^\circ\text{C}$) for Sonoran wrap assembly.
  * `low_watt_simmer_plate.hw`: $65\text{ W} - 140\text{ W}$ PID-controlled warming plate for tortillas, sauces, and reductions.
  * `sous_vide_stick.hw`: $180\text{ W}$ compact immersion circulator for field or counter prep.

### 03 · Prep & Actuation (`03_prep_actuation/`)
* **What it does:** Handles mechanical chopping, grinding, and portion dispensing on $12\text{ V} - 20\text{ V}$ rails.
* **Safety Guard:** Hall-effect lid interlocks (zero spin unless locked) + instant stall-current cutoff.

### 04 · Telemetry & Safety (`04_telemetry_safety/`)
* **What it does:** Continuously monitors heat, current draw, and session runtime.
* **Fail-Safes:**
  * **Hardware Watchdog:** Cuts gate drive within $15\text{ ms}$ if MCU firmware hangs.
  * **Dual NTC Delta Check:** Trips fault if primary and backup thermistors diverge by $>5^\circ\text{C}$.
  * **Session Auto-Off:** Configurable $30\text{ min}$ / $60\text{ min}$ idle shutoff.

---

## 🛠️ Quick-Start Checklist
- [ ] **Verify Cable Rating:** Use an E-Marked $100\text{ W}$ ($5\text{ A}$) or $240\text{ W}$ EPR USB-C cable.
- [ ] **Select Profile:** Set DIP switch or JSON config for target appliance (`SIMMER`, `BRIE_MELT`, `MILL`, `SOUS_VIDE`).
- [ ] **Confirm Telemetry:** Check serial/BLE readout for `PD_CONTRACT_OK` and `NTC_READY` before energizing load.

---

## 📄 License & Attribution
* **Framework:** Antistatic Lab Brand & Hardware Identity System
* **Research & Documentation:** Ejected Media / `localSq Research`
* **Hardware & Firmware:** Open-Source Hardware (CERN-OHL-P v2) / MIT