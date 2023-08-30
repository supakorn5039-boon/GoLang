import User from "../models/User.js";


//Update user and can change everything 
export const updateUser = async (req,res,next) => {
    try{
        const updatedUser = await User.findByIdAndUpdate(
            req.params.id,
            { $set: req.body},
            { new: true }
            )
        res.status(200).json(updatedUser)
    }catch(err){
        next(err)
    }
}

//Delete the user w/ ID
export const deleteUser = async (req,res,next) => {
    try{
        await User.findByIdAndDelete(
          req.params.id
          )
        res.status(200).json("User has been deleted.")
    }catch(err){
        next(err)
    }
}

//Find the User w/ ID 
export const getUser = async (req,res,next) => {
    try{
        const user = await User.findById(
            req.params.id
            )
        res.status(200).json(user)
    }catch(err){
        next(err)
    }
}

//Find all user Living in the Hotel w/ ID 
export const getUsers = async (req,res,next) => {
    try{
        const users = await User.find()
        res.status(200).json(users)
    }catch(err){
        next(err)
    }
}