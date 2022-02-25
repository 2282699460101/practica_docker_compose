import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from './home';
import Operacion from './operacion';
const Webpages = () => {
    return(
        <BrowserRouter>
            <Routes>
                <Route exact path="/" element = {<Home />} />
                <Route path = "/operacion" element = {<Operacion />} />
            </Routes>
        </BrowserRouter>
        
    );
};
export default Webpages;