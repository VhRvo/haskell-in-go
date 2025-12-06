{-# LANGUAGE ForeignFunctionInterface #-}

module Magic where

import Foreign.C.Types

-- 真正的 Haskell 函数
hsAdd :: CInt -> CInt -> IO CInt
hsAdd x y = pure (x + y + 121)

-- 把 hsAdd 导出成一个 C 符号 "add_hs"
foreign export ccall "add_hs"
  hsAdd :: CInt -> CInt -> IO CInt
