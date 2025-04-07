import express from 'express';
import cors from 'cors';
import rateLimit from 'express-rate-limit';

import routes from './routes/index';

const app = express();

const limiter = rateLimit({
  windowMs: 60 * 1000,
  max: 100,
  message: "Too many requests from this IP, please try again later.",
});

app.use(cors());
app.use(express.json());
app.use(limiter);

app.use("/api", routes);

export default app;
