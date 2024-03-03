import React, {useState} from 'react';
import { useNavigate } from "react-router-dom";

import { Tabs, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import PageTitle from '../components/pagetitle';
import { useAuth } from '../AuthContext';

import {Link} from 'react-router-dom'

Register.propTypes = {
};

function Register(props) {
    const [formData, setFormData] = useState({
        email: '',
        nickname: '',
        password: '',
        country: ''
    });

    const [errors, setErrors] = useState({});
    const navigate = useNavigate();
    const { login } = useAuth();

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.id]: e.target.value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        const errors = {};
        const passwordRegex = /^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-zA-Z]).{8,}$/;
        if (!formData.email.includes('@')) {
            errors.email = 'Please enter a valid email address.';
        }
        if (!passwordRegex.test(formData.password)) {
            errors.password = 'Password must be at least 8 characters long and include numbers and special characters.';
        }
        if (formData.password !== formData.password2) {
            errors.password2 = 'Passwords do not match.';
        }
        if (formData.nickname.trim() === '') {
            errors.nickname = 'Nickname is required.';
        }
        if (formData.country === '') {
            errors.country = 'Please select a country.';
        }


        if (Object.keys(errors).length === 0) {
            localStorage.setItem("token", "123");
            login()
        
            try {
                const response = await fetch("/api/user/register", {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(formData)
                });
    
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
    
                navigate("/");
            } catch (error) {
                console.error(error);
            }
        }
            
        setErrors(errors);
    };

    return (
        <div>
            <PageTitle heading='Register' title='Register' />

            <section className="register">
            <div className="container">
                <div className="row">
                <div className="col-md-12">
                    <div className="block-text center">
                    <h3 className="heading">Register To Rockie</h3>
                    <p className="desc fs-20">
                        Register in advance and enjoy the event benefits
                    </p>
                    </div>
                </div>
                <div className="col-md-12">
                <Tabs>
                    <TabPanel>
                        <div className="content-inner">
                            <form onSubmit={handleSubmit}>
                                <div className="form-group">
                                <label for="email">Email</label>
                                <input
                                    type="email"
                                    className="form-control"
                                    id="email"
                                    placeholder="Please fill in the email form."
                                    onChange={handleChange}
                                    autoComplete="username"
                                    required
                                />
                                {errors.email && <p className="error">{errors.email}</p>}
                                </div>
                                <div className="form-group">
                                <label
                                    >Password
                                    <span
                                    >(8 or more characters, including numbers and special
                                    characters)</span
                                    ></label
                                >
                                {errors.password && <p className="error">{errors.password}</p>}
                                <input
                                    type="password"
                                    className="form-control mb-10"
                                    id="password"
                                    placeholder="Please enter a password."
                                    onChange={handleChange}
                                    autoComplete="new-password"
                                    required
                                />
                                <input
                                    type="password"
                                    className="form-control"
                                    id="password2"
                                    placeholder="Please re-enter your password."
                                    onChange={handleChange}
                                    autoComplete="new-password"
                                    required
                                />
                                {errors.password2 && <p className="error">{errors.password2}</p>}
                                </div>
                                <div className="form-group">
                                <label for="nickname"
                                    >NickName
                                    <span className="fs-14"
                                    >(Excluding special characters)</span
                                    ></label
                                >
                                <input
                                    type="text"
                                    className="form-control"
                                    id="nickname"
                                    placeholder="Enter nickname"
                                    onChange={handleChange}
                                    required
                                />
                                {errors.nickname && <p className="error">{errors.nickname}</p>}
                                </div>
                                <div className="form-group">
                                <label for="country">Country </label>
                                <select className="form-control" onChange={handleChange} id="country" required>
                                    <option selected value="">Choose country</option>
                                    <option>Russia Federation</option>
                                    <option>USA</option>
                                    <option>South Korea</option>
                                </select>
                                {errors.country && <p className="error">{errors.country}</p>}
                                </div>
                                <button type="submit" className="btn-action">
                                Registration
                                </button>
                                <div className="bottom">
                                <p>Already have an account?</p>
                                <Link to="/login">Login</Link>
                                </div>
                            </form>
                        </div>
                    </TabPanel>
                </Tabs> 

                </div>
                </div>
            </div>
            </section>
        </div>
    );
}

export default Register;