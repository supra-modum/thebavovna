import React from "react";

interface AuthContextType {
  user: any;
  login: (callback: VoidFunction, token?: string) => void;
  logout: (callback: VoidFunction) => void;
}

export const AuthContext = React.createContext<AuthContextType>(null);