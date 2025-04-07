import dotenv from 'dotenv';
dotenv.config();

import express from 'express';
import app from './app';

const PORT = process.env.PORT || 3000;

app.listen(PORT, () => {
  console.log(`🚀 API Gateway running on port ${PORT}`);
});
