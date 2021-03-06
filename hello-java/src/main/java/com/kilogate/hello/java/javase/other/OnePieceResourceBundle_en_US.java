package com.kilogate.hello.java.javase.other;

import com.kilogate.hello.java.javase.core.basic.OnePieceRoleEnum;

import java.util.ListResourceBundle;

/**
 * 自定义的资源绑定类
 *
 * @author fengquanwei
 * @create 2017/8/4 18:08
 **/
public class OnePieceResourceBundle_en_US extends ListResourceBundle {
    private static final Object[][] contents = new Object[][]{
            {"Luffy", OnePieceRoleEnum.LUFFY},
            {"LuffyName", OnePieceRoleEnum.LUFFY.getEnglishhName()},
            {"Nami", OnePieceRoleEnum.NAMI},
            {"NamiName", OnePieceRoleEnum.NAMI.getEnglishhName()}
    };

    @Override
    protected Object[][] getContents() {
        return contents;
    }
}
