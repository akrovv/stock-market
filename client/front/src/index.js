import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import './App.scss'
import ScrollToTop from "./ScrollToTop";
import { AuthProvider } from './AuthContext';

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
    <React.Fragment>
        <AuthProvider>
        <BrowserRouter>
        <ScrollToTop />
            <App />
        </BrowserRouter>
        </AuthProvider>
    </React.Fragment>
);

