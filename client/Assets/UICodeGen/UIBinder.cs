﻿using UnityEngine;

public enum CodeGenObjectType
                if (trans.gameObject.GetComponent<UIBinder>() == null)
                {
                    trans.gameObject.AddComponent<UIBinder>();
                }

        if (go.GetComponent<Dropdown>() != null)
        {
            return CodeGenObjectType.GenAsDropdown;
        }