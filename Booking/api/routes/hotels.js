import express from "express";
import Hotel  from "../models/Hotel.js";
import { createHotel, deleteHotel, getHotel, getHotels, updateHotel } from "../controllers/hotel.js";
import { verifyAdmin } from "../utils/verifyToken.js";

const router = express.Router();

//create
router.post("/" ,verifyAdmin, createHotel);
//update
router.put("/:id" ,verifyAdmin, updateHotel);
//delete
router.delete("/:id" ,verifyAdmin,deleteHotel)
//get
router.get("/:id" ,verifyAdmin, getHotel)
//get all
router.get("/" , getHotels)

export default router