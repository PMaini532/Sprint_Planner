// src/pages/RegisterPage.js
import React, { useState } from "react";
import { registerUser } from "../api";
import { Link, useNavigate } from "react-router-dom";

export default function RegisterPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    const res = await registerUser(email, password);
    if (res.message) {
      setMessage("Registered successfully! You can now log in.");
      setTimeout(() => navigate("/login"), 1500);
    } else {
      setMessage(res.error || "Registration failed");
    }
  };

  return (
    <div className="auth-box">
      <h2>Register</h2>
      <form onSubmit={handleRegister}>
        <input type="email" placeholder="Email" value={email}
               onChange={(e) => setEmail(e.target.value)} required />
        <input type="password" placeholder="Password" value={password}
               onChange={(e) => setPassword(e.target.value)} required />
        <button type="submit">Register</button>
        <p className="message">{message}</p>
        <p>Already have an account? <Link to="/login">Login</Link></p>
      </form>
    </div>
  );
}
