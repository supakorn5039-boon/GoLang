import mongoose from 'mongoose'


//What User have If They Register
const UserSchema = new mongoose.Schema({
    username:{
        type:String,
        require:true,
        unique:true
    },
    email:{
        type:String,
        require:true,
        unique:true
    },
    password:{
        type:String,
        require:true
    },
    isAdmin:{
        type:Boolean,
        default:false,
    },
  },
  { timestamps: true }

);

export default mongoose.model("User" , UserSchema)