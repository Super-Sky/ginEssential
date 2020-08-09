// const { axios } = require("vue/types/umd");

import axios from 'axios';
import storageService from '../service/storageService';

const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    timeout: 1000,
    headers: { Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}` },
});

export default service;
