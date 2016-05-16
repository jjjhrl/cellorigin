﻿// Generated by github.com/davyxu/cellorigin
using UnityEngine;
using UnityEngine.UI;

partial class LoginCharInfoView : Framework.BaseView
{
	LoginCharInfoPresenter _Presenter;
	
	Button _SelectChar;
	Text _CharName;
	
	public override void Bind( Framework.BasePresenter presenter )
	{
		_Presenter = presenter as LoginCharInfoPresenter;
		
		var trans = this.transform;
		
		_SelectChar = trans.Find("SelectChar").GetComponent<Button>();
		_CharName = trans.Find("SelectChar/CharName").GetComponent<Text>();
		
		_SelectChar.onClick.AddListener( _Presenter.Cmd_SelectChar );
		_Presenter.OnCharNameChanged += delegate()
		{
			_CharName.text = _Presenter.CharName;
		};
		if ( _Presenter.CharName != null )
		{
			_CharName.text = _Presenter.CharName;
		}
		
	}
	
}