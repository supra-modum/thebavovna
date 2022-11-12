import { Routes, Route } from "react-router-dom";
import HomePage from "~src/pages/HomePage";
import AuthPage from "~src/pages/AuthPage";

export default function App() {
  return (
    <div>
      {/* <Routes> */}
      {/*   <Route path="/" element={<HomePage />} /> */}
      {/*   <Route path="/login" element={<AuthPage />} /> */}
      {/*   <Route path="*" element={<HomePage />} /> */}
      {/* </Routes> */}
      <HomePage/>
    </div>
  );
}
