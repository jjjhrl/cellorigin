﻿//this source code was auto-generated by tolua#, do not modify it
using System;
using LuaInterface;

public class DescriptorWrap
{
	public static void Register(LuaState L)
	{
		L.BeginClass(typeof(Descriptor), typeof(ProtoBase));
		L.RegFunction("GetFieldByName", GetFieldByName);
		L.RegFunction("GetFieldByFieldNumber", GetFieldByFieldNumber);
		L.RegFunction("GetNestedMessage", GetNestedMessage);
		L.RegFunction("GetNestedEnum", GetNestedEnum);
		L.RegFunction("__tostring", Lua_ToString);
		L.RegVar("Define", get_Define, null);
		L.RegVar("Name", get_Name, null);
		L.EndClass();
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int GetFieldByName(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 2);
			Descriptor obj = (Descriptor)ToLua.CheckObject(L, 1, typeof(Descriptor));
			string arg0 = ToLua.CheckString(L, 2);
			FieldDescriptor o = obj.GetFieldByName(arg0);
			ToLua.PushObject(L, o);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int GetFieldByFieldNumber(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 2);
			Descriptor obj = (Descriptor)ToLua.CheckObject(L, 1, typeof(Descriptor));
			int arg0 = (int)LuaDLL.luaL_checknumber(L, 2);
			FieldDescriptor o = obj.GetFieldByFieldNumber(arg0);
			ToLua.PushObject(L, o);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int GetNestedMessage(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 2);
			Descriptor obj = (Descriptor)ToLua.CheckObject(L, 1, typeof(Descriptor));
			string arg0 = ToLua.CheckString(L, 2);
			Descriptor o = obj.GetNestedMessage(arg0);
			ToLua.PushObject(L, o);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int GetNestedEnum(IntPtr L)
	{
		try
		{
			ToLua.CheckArgsCount(L, 2);
			Descriptor obj = (Descriptor)ToLua.CheckObject(L, 1, typeof(Descriptor));
			string arg0 = ToLua.CheckString(L, 2);
			EnumDescriptor o = obj.GetNestedEnum(arg0);
			ToLua.PushObject(L, o);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int Lua_ToString(IntPtr L)
	{
		object obj = ToLua.ToObject(L, 1);

		if (obj != null)
		{
			LuaDLL.lua_pushstring(L, obj.ToString());
		}
		else
		{
			LuaDLL.lua_pushnil(L);
		}

		return 1;
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get_Define(IntPtr L)
	{
		object o = null;

		try
		{
			o = ToLua.ToObject(L, 1);
			Descriptor obj = (Descriptor)o;
			google.protobuf.DescriptorProto ret = obj.Define;
			ToLua.PushObject(L, ret);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e, o == null ? "attempt to index Define on a nil value" : e.Message);
		}
	}

	[MonoPInvokeCallbackAttribute(typeof(LuaCSFunction))]
	static int get_Name(IntPtr L)
	{
		object o = null;

		try
		{
			o = ToLua.ToObject(L, 1);
			Descriptor obj = (Descriptor)o;
			string ret = obj.Name;
			LuaDLL.lua_pushstring(L, ret);
			return 1;
		}
		catch(Exception e)
		{
			return LuaDLL.toluaL_exception(L, e, o == null ? "attempt to index Name on a nil value" : e.Message);
		}
	}
}

