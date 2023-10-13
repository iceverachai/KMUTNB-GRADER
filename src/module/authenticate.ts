import { useUserStore } from "./userstore"

export const authenticate = async () : Promise<boolean> => {
    const userstore = useUserStore()
    try{
        const token = localStorage.getItem('token');
        if(!token?.length) throw new Error('Token empty or undefined');

        const response = await fetch(import.meta.env.VITE_api+"/api/accessgate", {
            method: "GET",
            headers:{
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`
            }
        });
        const json = await response.json();
        if(json.status === "success" && json.response){
            userstore.setInfo(json.response);
            return true;
        }
        return false;
    }catch{
        userstore.setInfo(null);
        return false;
    }
}

export default authenticate;