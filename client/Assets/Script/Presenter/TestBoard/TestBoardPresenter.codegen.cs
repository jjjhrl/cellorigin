﻿// Generated by github.com/davyxu/cellorigin
using UnityEngine;
using UnityEngine.UI;
using System;


partial class TestBoardPresenter : Framework.BasePresenter
{
	TestBoardModel _Model;

    public Action OnTextInfoChanged;
    public virtual string TextInfo
    {
        get
        {
            return _Model.TextInfo;
        }
        set
        {
            _Model.TextInfo = value;

            if (OnTextInfoChanged != null)
            {
                OnTextInfoChanged();
            }
        }
    }


	
	public void Init( )
	{
		_Model = Framework.ModelManager.Instance.Get<TestBoardModel>();
		
	}
	
}
