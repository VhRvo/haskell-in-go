module Main where

import qualified Magic (hsAdd)

main :: IO ()
main = do
  putStrLn "Hello, Haskell!"
  putStrLn "Calling hsAdd(1, 2):"
  print =<< Magic.hsAdd 1 2
