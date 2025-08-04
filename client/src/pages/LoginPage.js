import React, { useState } from "react";
import { loginUser } from "../api";
import { useNavigate, Link } from "react-router-dom";
import { jwtDecode } from "jwt-decode";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    const res = await loginUser(email, password);
    if (res.token) {
      localStorage.setItem("token", res.token);
      const decoded = jwtDecode(res.token);
      setMessage("Login successful!");
      if (decoded.role === "admin") {
      navigate("/admin");  // Replace with your actual admin route
      }else {
        navigate("/dashboard"); // Developer route
      }
      } else {
        setMessage(res.error || "Login failed");
      }
  };

  return (
    <div className="auth-box">
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <input type="email" placeholder="Email" value={email}
               onChange={(e) => setEmail(e.target.value)} required />
        <input type="password" placeholder="Password" value={password}
               onChange={(e) => setPassword(e.target.value)} required />
        <button type="submit">Login</button>
        <p className="message">{message}</p>
        <p>Don't have an account? <Link to="/register">Register</Link></p>
      </form>
    </div>
  );
}
